package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"redpacket/config"
)

var (
	DB     *gorm.DB
	Config = &config.Config{}
	LOG    *zap.Logger
	Redis  *redis.Client
)
