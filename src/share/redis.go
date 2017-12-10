package share

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func NewRedisPool(maxIdle, maxActive, DBNum int, timeout time.Duration, addr, password string) *redis.Pool {

	return &redis.Pool{
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
		IdleTimeout: timeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			// return redis.DialURL(rawurl)
			// return redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(dbNum))
			return redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(DBNum))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
