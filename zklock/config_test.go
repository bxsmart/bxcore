package zklock_test

import (
	"github.com/bxsmart/bxcore/log"
	"github.com/bxsmart/bxcore/zklock"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	log.Initialize()
	config := zklock.ZkLockConfig{}
	config.ZkServers = "127.0.0.1:2181"
	config.ConnectTimeOut = 2
	zklock.Initialize(config)
	zklock.RegisterConfigHandler("demoService", "name", func(data []byte) error {
		str := string(data[:])
		log.Infof("data changed, new data is : %s", str)
		return nil
	})
	time.Sleep(1000 * time.Second)
}
