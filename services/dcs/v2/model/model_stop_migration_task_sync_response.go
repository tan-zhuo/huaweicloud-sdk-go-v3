package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopMigrationTaskSyncResponse Response Object
type StopMigrationTaskSyncResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMigrationTaskSyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncResponse struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncResponse", string(data)}, " ")
}
