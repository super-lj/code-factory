package config

import (
	"ci/logs"
	"ci/manager"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() {
	initDB()
	initRedis()
}

func initRedis() {
	opt := &redis.Options{}
	manager.CodeFactoryRedis = redis.NewClient(opt)
}

func initDB() {
	ctx := context.Background()
	db, err := gorm.Open("sqlite3", "code_factory_ci.db")
	if err != nil {
		logs.CtxError(ctx, fmt.Sprintf("open db fail, err: %v", err))
		panic(err)
	}
	manager.CodeFactoryDBRead = db
}
