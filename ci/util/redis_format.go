package util

import (
	"fmt"
)

const (
	CommitRedisKey = "commit:v1:%s"
)

func GetCommitRedisKey(hash string) string {
	return fmt.Sprintf(CommitRedisKey, hash)
}
