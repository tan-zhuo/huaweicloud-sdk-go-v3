package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddMembersRequest Request Object
type BatchAddMembersRequest struct {
	Body *BatchAddMembersRequestBody `json:"body,omitempty"`
}

func (o BatchAddMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchAddMembersRequest", string(data)}, " ")
}
