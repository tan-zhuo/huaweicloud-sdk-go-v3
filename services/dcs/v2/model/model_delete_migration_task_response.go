package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigrationTaskResponse Response Object
type DeleteMigrationTaskResponse struct {

	// 删除的迁移任务ID列表。
	TaskIdList     *[]string `json:"task_id_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigrationTaskResponse", string(data)}, " ")
}
