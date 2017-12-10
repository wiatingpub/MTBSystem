package main

import (
	"domain/apid"
	"encoding/json"
	"github.com/micro/go-micro/cmd"
	microErrors "github.com/micro/go-micro/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"share/config"
	"share/utils/log"
	"strconv"
)

var (
	cors   = map[string]bool{"*": true}
	logger *zap.Logger
)

func init() {
	logger = log.Init("api")

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRPC)
	http.ListenAndServe(":8082", mux)
	logger.Info("Listen on :8082")
}

func handleRPC(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("ok,this is the server ..."))
		return
	}

	// 跨域处理
	if origin := r.Header.Get("Origin"); cors[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else if len(origin) > 0 && cors["*"] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handleJSONRPC(w, r)
	return
}

func handleJSONRPC(w http.ResponseWriter, r *http.Request) {

	service, method := apid.PathToReceiver(config.Namespace, r.URL.Path)

	br, _ := ioutil.ReadAll(r.Body)

	request := json.RawMessage(br)

	var response json.RawMessage
	req := (*cmd.DefaultOptions().Client).NewJsonRequest(service, method, &request)
	ctx := apid.RequestToContext(r)
	err := (*cmd.DefaultOptions().Client).Call(ctx, req, &response)
	// make the call
	if err != nil {
		ce := microErrors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = service
			ce.Status = http.StatusText(500)
			// ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}
	b, _ := response.MarshalJSON()
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}
