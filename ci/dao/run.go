package dao

import "gorm.io/gorm"

type Run struct {
	gorm.Model
	RepoName   string
	BranchName string
	CommitHash string
	Num        int32
	Status     string
	Log        string
}
