package domain

import (
	"ci/enumeration"
	"github.com/jinzhu/gorm"
)

type Repo struct {
	gorm.Model
	Id      int64                      `json:"id"`
	Name    string                     `json:"name"`
	Url     string                     `json:"url"`
	Status  enumeration.RepoStatusEnum `json:"status"`
	History []*Commit                  `json:"history"`
}
