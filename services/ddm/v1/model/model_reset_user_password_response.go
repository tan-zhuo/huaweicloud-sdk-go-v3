package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserPasswordResponse Response Object
type ResetUserPasswordResponse struct {

	// 操作是否成功。
	Success *bool `json:"success,omitempty"`

	// DDM实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// DDM账号名称
	UserName       *string `json:"user_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordResponse", string(data)}, " ")
}
