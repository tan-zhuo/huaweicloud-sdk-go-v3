package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogGroupResponse Response Object
type CreateLogGroupResponse struct {

	// 创建的日志组的Id。
	LogGroupId     *string `json:"log_group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateLogGroupResponse", string(data)}, " ")
}
