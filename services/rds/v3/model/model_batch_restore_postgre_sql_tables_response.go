package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestorePostgreSqlTablesResponse Response Object
type BatchRestorePostgreSqlTablesResponse struct {

	// 表信息
	RestoreResult  *[]RestoreResult `json:"restore_result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchRestorePostgreSqlTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestorePostgreSqlTablesResponse struct{}"
	}

	return strings.Join([]string{"BatchRestorePostgreSqlTablesResponse", string(data)}, " ")
}
