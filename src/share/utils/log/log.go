package log

import (
	"bytes"
	"log"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"share/config"
)

var instance *zap.Logger

// Instance 唯一实例
func Instance() *zap.Logger {
	return instance
}

// Init 初始化
func Init(serviceName string) *zap.Logger {
	instance = NewLogger(serviceName)
	return instance
}

// NewLogger 新建日志
func NewLogger(srvName string) *zap.Logger {

	directory := path.Join(config.LogPath, srvName)
	writers := []zapcore.WriteSyncer{newRollingFile(directory)}
	writers = append(writers, os.Stdout)
	logger, dyn := newZapLogger(true, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			updateLogLevel(srvName, dyn, false)
		}
	}()

	return logger
}

func updateLogLevel(serviceName string, dyn *zap.AtomicLevel, isProduction bool) {

	originLevelString := "info"
	if !isProduction {
		originLevelString = "debug"
	}

	levelConf := make(map[string]map[string]string)

	newLevelString, ok := levelConf[serviceName]["127.0.0.1"]
	if !ok {
		newLevelString, ok = levelConf[serviceName]["*"]
		if !ok {
			newLevelString = originLevelString
		}
	}

	if !ok {
		newLevelString, ok = levelConf[serviceName]["*"]
		if !ok {
			newLevelString = originLevelString
		}
	}

	newLevel := new(zapcore.Level)
	if err := newLevel.Set(newLevelString); err != nil {
		newLevel.Set(originLevelString)
	}
	if dyn.Level() != *newLevel {
		log.Println("修改日志等级: ", dyn.Level().String(), "=>", newLevel.String())
		dyn.SetLevel(*newLevel)
	}
}

func newRollingFile(directory string) zapcore.WriteSyncer {
	if err := os.MkdirAll(directory, 0766); err != nil {
		log.Println("failed create log directory:", directory, ":", err)
		return nil
	}

	return newLumberjackWriteSyncer(&lumberjack.Logger{
		Filename:  path.Join(directory, "output.log"),
		MaxSize:   100, //megabytes
		MaxAge:    7,   //days
		LocalTime: true,
		Compress:  false,
	})
}

func newZapLogger(isProduction bool, output zapcore.WriteSyncer) (*zap.Logger, *zap.AtomicLevel) {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
	}

	var encoder zapcore.Encoder
	dyn := zap.NewAtomicLevel()
	if isProduction {
		dyn.SetLevel(zap.InfoLevel)
		encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encCfg) // zapcore.NewJSONEncoder(encCfg)
	} else {
		dyn.SetLevel(zap.DebugLevel)
		encCfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}

	return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller()), &dyn
}

type lumberjackWriteSyncer struct {
	*lumberjack.Logger
	buf       *bytes.Buffer
	logChan   chan []byte
	closeChan chan interface{}
	maxSize   int
}

func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
	ws := &lumberjackWriteSyncer{
		Logger:    l,
		buf:       bytes.NewBuffer([]byte{}),
		logChan:   make(chan []byte, 5000),
		closeChan: make(chan interface{}),
		maxSize:   1024,
	}
	go ws.run()
	return ws
}

func (l *lumberjackWriteSyncer) run() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if l.buf.Len() > 0 {
				l.sync()
			}
		case bs := <-l.logChan:
			_, err := l.buf.Write(bs)
			if err != nil {
				continue
			}
			if l.buf.Len() > l.maxSize {
				l.sync()
			}
		case <-l.closeChan:
			l.sync()
			return
		}
	}
}

func (l *lumberjackWriteSyncer) Stop() {
	close(l.closeChan)
}

func (l *lumberjackWriteSyncer) Write(bs []byte) (int, error) {
	b := make([]byte, len(bs))
	for i, c := range bs {
		b[i] = c
	}
	l.logChan <- b
	return 0, nil
}

func (l *lumberjackWriteSyncer) Sync() error {
	return nil
}

func (l *lumberjackWriteSyncer) sync() error {
	defer l.buf.Reset()
	_, err := l.Logger.Write(l.buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
