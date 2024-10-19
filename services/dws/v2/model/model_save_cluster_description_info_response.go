package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveClusterDescriptionInfoResponse Response Object
type SaveClusterDescriptionInfoResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveClusterDescriptionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveClusterDescriptionInfoResponse struct{}"
	}

	return strings.Join([]string{"SaveClusterDescriptionInfoResponse", string(data)}, " ")
}
