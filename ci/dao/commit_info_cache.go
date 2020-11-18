package dao

import (
	"ci/domain"
	"ci/manager"
	"ci/util"
	"encoding/json"
	"fmt"
	"time"
)

// GetCommitByHashFromCache returns a commit associated with given commit id
func GetCommitByHashFromCache(hash string) (*domain.Commit, util.OpResult) {
	key := util.GetCommitRedisKey(hash)
	res, err := manager.CodeFactoryRedis.Get(key).Result()
	if err != nil {
		return nil, util.NewOpResult(util.ErrRedis, fmt.Sprintf("get commit from redis err, hash: %s", hash))
	}
	var commit *domain.Commit
	err = json.Unmarshal([]byte(res), commit)
	if err != nil {
		return nil, util.NewOpResult(util.ErrSystemInternal, fmt.Sprintf("unmarshal from string err"))
	}
	return commit, util.NewSucOpResult()
}

func SaveCommitToCache(commit *domain.Commit) util.OpResult {
	key := util.GetCommitRedisKey(commit.Hash)
	b, err := json.Marshal(commit)
	if err != nil {
		return util.OpResult{util.ErrSystemInternal, fmt.Sprintf("json marshal fail, err: %+v", err)}
	}
	duration := time.Hour * 24 * 7
	_, err = manager.CodeFactoryRedis.Set(key, b, duration).Result()
	if err != nil {
		return util.NewOpResult(util.ErrRedis, fmt.Sprintf("save commit from redis err, hash: %s", commit.Hash))
	}
	return util.NewSucOpResult()
}
