package dao

import (
	"ci/domain"
	"ci/enumeration"
	"ci/manager"
	"ci/util"
	"github.com/jinzhu/gorm"
	"context"
	"fmt"
)

func GetRepoInfoByHash(ctx context.Context, hash string) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	err := manager.CodeFactoryDBRead.Where(fmt.Sprintf("hash = %d", hash)).First(&repo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NewSucOpResult()
		} else {
			return nil, util.NewOpResult(util.ErrDBRead, "[GetRepoInfoById] get db record err")
		}
	}
	return repo, util.NewSucOpResult()
}

func GetRepoInfoByStatus(ctx context.Context, status enumeration.RepoStatusEnum) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	err := manager.CodeFactoryDBRead.Where(fmt.Sprintf("status = %d", status)).First(&repo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NewSucOpResult()
		} else {
			return nil, util.NewOpResult(util.ErrDBRead, "[GetRepoInfoByStatus] get db record err")
		}
	}
	return repo, util.NewSucOpResult()
}
