package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAclRulesResponse Response Object
type BatchDeleteAclRulesResponse struct {
	Data           *BatchDeleteAclRulesResponseData `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o BatchDeleteAclRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAclRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclRulesResponse", string(data)}, " ")
}
