package test

import (
	"fmt"
	"github.com/kgip/redis-lock/adapters"
	"github.com/kgip/redis-lock/lock"
	"redpacket/global"
	"redpacket/initialize"
	"redpacket/model/po"
	"redpacket/service/impl"
	"testing"
	"time"
)

var userService = impl.UserService{}

func init() {
	////1.初始化配置文件ls
	initialize.Config(fmt.Sprintf("../%s", global.ConfigPath))
	////2.初始化zap日志
	global.LOG = initialize.Zap()
	//
	global.DB = initialize.Gorm()

	global.Redis = initialize.Redis()
}

func TestAddUser(t *testing.T) {
	//users := []*po.User{{}, {}}
	//userService.AddUsers(users)
	//var balance int64
	user := &po.User{}
	global.DB.Model(&po.User{}).Select("balance", "username").Where("username in (?)", []string{"aaa", "bbb"}).Limit(1).Scan(&user)
	t.Log(user)
}

func TestRedisLock(t *testing.T) {
	lockOperator := lock.NewRedisLockOperator(adapters.NewGoRedisV8Adapter(global.Redis), lock.EnableWatchdog(true))
	key := "key1"
	lockOperator.Lock(key, lock.Context())
	for i := 0; i < 10; i++ {
		t.Log("lock test!!!")
		time.Sleep(3 * time.Second)
	}
	lockOperator.Unlock(key)
}
