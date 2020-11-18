package manager

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	CodeFactoryDBRead *gorm.DB
	CodeFactoryRedis  *redis.Client
)
