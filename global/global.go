package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/kgip/redis-lock/lock"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"redpacket/config"
)

var (
	DB           *gorm.DB
	Config       = &config.Config{}
	LOG          *zap.Logger
	Redis        *redis.Client
	LockOperator *lock.RedisLockOperator //redis分布式锁
)
