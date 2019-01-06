package zklock

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"sync"
	"time"
)

/*todo:
1、通过zk设置配置文件
2、节点加入与退出事件
*/

type ZkLockConfig struct {
	ZkServers      string
	ConnectTimeOut int
}

type ZkLock struct {
	lockMap map[string]*zk.Lock
	mutex   sync.Mutex
}

var ZkClient *zk.Conn
var zl *ZkLock

const lockBasePath = "/loopring_lock"

func Initialize(config ZkLockConfig) (*ZkLock, error) {
	if config.ZkServers == "" || len(config.ZkServers) < 10 {
		return nil, fmt.Errorf("Zookeeper server list config invalid: %s\n", config.ZkServers)
	}
	var err error
	ZkClient, _, err = zk.Connect(strings.Split(config.ZkServers, ","), time.Second*time.Duration(config.ConnectTimeOut))
	if err != nil {
		return nil, fmt.Errorf("Connect zookeeper error: %s\n", err.Error())
	}
	zl = &ZkLock{make(map[string]*zk.Lock), sync.Mutex{}}
	return zl, nil
}

//when get err, should send sns message
func TryLock(lockName string) error {
	zl.mutex.Lock()
	if _, ok := zl.lockMap[lockName]; !ok {
		acls := zk.WorldACL(zk.PermAll)
		zl.lockMap[lockName] = zk.NewLock(ZkClient, fmt.Sprintf("%s/%s", lockBasePath, lockName), acls)
	}
	zl.mutex.Unlock()
	return zl.lockMap[lockName].Lock()
}

func ReleaseLock(lockName string) error {
	if innerLock, ok := zl.lockMap[lockName]; ok {
		innerLock.Unlock()
		return nil
	} else {
		return fmt.Errorf("Try release not exists lock: %s\n", lockName)
	}
}

func IsLockInitialed() bool {
	return nil != zl
}


func CreatePath(path string) (string, error) {
	isExist, _, err := ZkClient.Exists(path)
	if err != nil {
		return "", fmt.Errorf("check zk path %s exists failed : %s", path, err.Error())
	}
	if !isExist {
		_, err := ZkClient.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
		if err != nil && err != zk.ErrNodeExists {
			return "", fmt.Errorf("failed create balancer sub path %s ,with error : %s ", path, err.Error())
		}
	}
	return path, nil
}