package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicAttributesResponse Response Object
type DeleteTopicAttributesResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributesResponse", string(data)}, " ")
}
