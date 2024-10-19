package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationResponse Response Object
type UpdateApplicationResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationResponse", string(data)}, " ")
}
