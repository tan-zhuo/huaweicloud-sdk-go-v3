package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeployKeyResponse Response Object
type AddDeployKeyResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *Key `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDeployKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyResponse struct{}"
	}

	return strings.Join([]string{"AddDeployKeyResponse", string(data)}, " ")
}
