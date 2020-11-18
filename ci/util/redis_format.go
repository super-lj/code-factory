package util

import (
	"fmt"
	"strconv"
)

const (
	CommitRedisKey = "commit:v1:%s"
)

func GetCommitRedisKey(id int64) string {
	idStr := strconv.Itoa(int(id))
	return fmt.Sprintf(CommitRedisKey, idStr)
}
