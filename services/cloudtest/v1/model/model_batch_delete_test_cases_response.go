package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestCasesResponse Response Object
type BatchDeleteTestCasesResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueString `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId *string `json:"request_id,omitempty"`

	// 对内接口才有此属性
	ServerAddress  *string `json:"server_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCasesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCasesResponse", string(data)}, " ")
}
