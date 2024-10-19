package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionInfoResponse Response Object
type CreateConnectionInfoResponse struct {

	// 状态码
	ResCode *int32 `json:"res_code,omitempty"`

	// 成功信息
	ResLog *string `json:"res_log,omitempty"`

	// 成功信息
	ResMsg         *string `json:"res_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConnectionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionInfoResponse", string(data)}, " ")
}
