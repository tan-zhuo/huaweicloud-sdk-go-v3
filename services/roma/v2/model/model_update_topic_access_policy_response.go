package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicAccessPolicyResponse Response Object
type UpdateTopicAccessPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyResponse", string(data)}, " ")
}
