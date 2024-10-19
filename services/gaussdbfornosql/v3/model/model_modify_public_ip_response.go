package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPublicIpResponse Response Object
type ModifyPublicIpResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicIpResponse struct{}"
	}

	return strings.Join([]string{"ModifyPublicIpResponse", string(data)}, " ")
}
