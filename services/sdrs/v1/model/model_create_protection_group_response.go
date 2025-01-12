package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProtectionGroupResponse Response Object
type CreateProtectionGroupResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupResponse", string(data)}, " ")
}
