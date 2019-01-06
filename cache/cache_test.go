package cache_test

import (
	"encoding/json"
	"github.com/bxsmart/bxcore/cache"
	"github.com/bxsmart/bxcore/cache/redis"
	"github.com/bxsmart/bxcore/log"
	"testing"
	"time"
)

func init() {
	log.Initialize()
	cache.NewCache(redis.RedisOptions{Host: "192.168.10.240", Port: "6379", IdleTimeout: 20, Database: 5, MaxActive: 20, MaxIdle: 20})
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func TestGet(t *testing.T) {
	cacheKey := "conn_test"
	for i := 1; i <= 2; i++ {
		user := &User{
			Id:   int64(i + 100),
			Name: "name1",
		}
		if data, err := json.Marshal(user); nil != err {
			log.Debugf("err:%s", err.Error())
		} else {
			go func(data []byte) {
				defer func() {
					if r := recover(); r != nil {
						println(r)
						log.Info("llllllllllll")
						//log.Errorf("Recovered in f", r)
						//log.Errorf("Recovered in f", r)
					}
				}()
				f(cacheKey, data)
				f1(cacheKey, data)
				f2()
				time.Sleep(5 * time.Second)
			}(data)
		}
	}

	time.Sleep(10 * time.Second)
}

func f(cacheKey string, data []byte) {
	var datalist [][]byte = [][]byte{[]byte("777"), []byte("123")}
	if err := cache.SAdd(cacheKey+"_0000", int64(10000), datalist...); nil != err {
		log.Errorf("err:%s", err.Error())
	}
}

func f1(cacheKey string, data []byte) {
	if err := cache.SAdd(cacheKey, int64(10000), data); nil != err {
		log.Errorf("err:%s", err.Error())
	}
}

func f2() {
	cache.Publish("aaa", []byte("1231321312"))
	panic("eeeeeeee")
}
