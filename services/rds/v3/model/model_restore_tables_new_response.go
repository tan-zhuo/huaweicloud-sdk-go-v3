package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreTablesNewResponse Response Object
type RestoreTablesNewResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreTablesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesNewResponse struct{}"
	}

	return strings.Join([]string{"RestoreTablesNewResponse", string(data)}, " ")
}
