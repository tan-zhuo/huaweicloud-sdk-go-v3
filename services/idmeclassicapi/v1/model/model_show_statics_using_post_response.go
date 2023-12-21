package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStaticsUsingPostResponse Response Object
type ShowStaticsUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]interface{} `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowStaticsUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStaticsUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowStaticsUsingPostResponse", string(data)}, " ")
}
