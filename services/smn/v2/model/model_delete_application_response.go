package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationResponse Response Object
type DeleteApplicationResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationResponse", string(data)}, " ")
}
