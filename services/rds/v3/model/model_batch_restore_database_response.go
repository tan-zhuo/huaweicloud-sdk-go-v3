package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestoreDatabaseResponse Response Object
type BatchRestoreDatabaseResponse struct {

	// 表信息
	RestoreResult  *[]RestoreResult `json:"restore_result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchRestoreDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreDatabaseResponse struct{}"
	}

	return strings.Join([]string{"BatchRestoreDatabaseResponse", string(data)}, " ")
}
