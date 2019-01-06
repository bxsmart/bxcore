package zklock

import (
	"fmt"
	"github.com/bxsmart/bxcore/log"
	"github.com/samuel/go-zookeeper/zk"
)

const configShareBasePath = "/bxsmart_config"

type HandlerFunc func(value []byte) error

func RegisterConfigHandler(namespace string, key string, action HandlerFunc) error {
	if !IsLockInitialed() {
		return fmt.Errorf("zkClient is not intiliazed")
	}
	if _, err := CreatePath(configShareBasePath); err != nil {
		return fmt.Errorf("create config base path failed %s with error %s", configShareBasePath, err.Error())
	}
	ns := fmt.Sprintf("%s/%s", configShareBasePath, namespace)
	if _, err := CreatePath(ns); err != nil {
		return fmt.Errorf("create config namespace path failed %s with error %s", ns, err.Error())
	}
	keyPath := fmt.Sprintf("%s/%s", ns, key)
	if _, err := CreatePath(keyPath); err != nil {
		return fmt.Errorf("create config key path failed %s with error %s", keyPath, err.Error())
	}

	if data, _, ch, err := ZkClient.GetW(keyPath); err != nil {
		log.Errorf("Get config %s data failed with error : %s", keyPath, err.Error())
	} else {
		log.Infof("Get config %s data success", keyPath)
		action(data)
		go func() {
			for {
				select {
				case evt := <-ch:
					if evt.Type == zk.EventNodeDataChanged {
						if data, _, chx, err := ZkClient.GetW(keyPath); err != nil {
							log.Errorf("Get config %s data failed with error : %s", keyPath, err.Error())
						} else {
							log.Infof("config %s data changed to value %s", keyPath, string(data[:]))
							ch = chx
							action(data)
						}
					}
				}
			}
		}()
	}
	return nil
}
