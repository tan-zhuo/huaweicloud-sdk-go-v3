package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeployKeyResponse Response Object
type DeleteDeployKeyResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *bool `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeployKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeployKeyResponse", string(data)}, " ")
}
