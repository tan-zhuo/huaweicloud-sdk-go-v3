package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUndoCheckoutUsingPostResponse Response Object
type ShowUndoCheckoutUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUndoCheckoutUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUndoCheckoutUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowUndoCheckoutUsingPostResponse", string(data)}, " ")
}
