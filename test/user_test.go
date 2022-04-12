package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kgip/redis-lock/adapters"
	"github.com/kgip/redis-lock/lock"
	"github.com/shopspring/decimal"
	"math/rand"
	"redpacket/ex"
	"redpacket/global"
	"redpacket/initialize"
	"redpacket/model/common"
	"redpacket/model/po"
	"redpacket/model/vo"
	"redpacket/service/impl"
	"redpacket/utils"
	"redpacket/utils/mq"
	"testing"
	"time"
)

var userService = impl.UserService{}

func init() {
	////1.初始化配置文件ls
	initialize.Config(fmt.Sprintf("../%s", global.ConfigPath))
	////2.初始化zap日志
	global.LOG = initialize.Zap()
	////
	//global.DB = initialize.Gorm()

	//global.Redis = initialize.Redis()
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

func TestLocalMQ(t *testing.T) {
	mqOperator := mq.NewLocalMQ().AddQueue("topic1", 1000)
	mqOperator.SendMessage("topic1", "aaa", 3*time.Second)
	mqOperator.RegistryMessageHandler([]string{"topic1"}, func(msg interface{}) {
		t.Log(msg)
	})
	time.Sleep(10 * time.Second)
}

func TestEx(t *testing.T) {
	//var e interface{} = ex.InternalException
	//if _, ok := e.(*ex.Exception); ok {
	//	t.Log("ok")
	//} else {
	//	t.Error("error")
	//}
	bytes, _ := json.Marshal(ex.InternalException)
	t.Log(string(bytes))
}

func TestCopy(t *testing.T) {
	user := &po.User{Base: po.Base{ID: 1, CreatedAt: time.Now()}, Username: "", Balance: 1111, Age: 18}
	userVo := &vo.UserVo{CreatedAt: common.JSONTime(user.CreatedAt)}
	utils.BeanCopy(user, userVo, "CreatedAt")
	t.Log(userVo)
}

func TestRedisDel(t *testing.T) {
	global.Redis.Set(context.Background(), "aaa", 1, redis.KeepTTL)
	t.Log(global.Redis.Del(context.Background(), "aaaa").Result())
}

func TestPtr(t *testing.T) {
	t.Logf("%p", ex.InternalException)
	t.Logf("%p", ex.InternalException.SetDetail("aaa"))
}

func TestRedisZSet(t *testing.T) {
	t.Log(global.Redis.ZRangeWithScores(context.Background(), "test", 0, -1).Result())
	result, err := global.Redis.ZRank(context.Background(), "test", "ccc").Result()
	t.Log(err)
	t.Log(result)
}

func TestRedisList(t *testing.T) {
	global.Redis.RPush(context.Background(), "testList", 111)
	for i := 0; i < 5; i++ {
		t.Log(global.Redis.LPop(context.Background(), "testList").Result())
	}
	global.Redis.Del(context.Background(), "testList")
	for i := 0; i < 5; i++ {
		t.Log(global.Redis.LPop(context.Background(), "testList").Result())
	}
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(10))
		t.Log(rand.Float64())
	}
}

func TestGenericRedPacket(t *testing.T) {
	for i := 0; i < 10; i++ {
		packets := utils.GenericRedPackets(13, 10)
		t.Log(packets)
		var total float64 = 0
		for _, balance := range packets {
			total += balance.(float64)
		}
		t.Log(total)
		time.Sleep(time.Second)
	}
}

func TestDecimal(t *testing.T) {
	f := decimal.NewFromFloat(1.0101)
	t.Log(f.Add(decimal.NewFromInt(1)).RoundDown(2).Float64())
	t.Log(f.Float64())
}
