package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberRequestBody This is a auto create Body Object
type UpdateMemberRequestBody struct {
	Member *UpdateMemberOption `json:"member"`
}

func (o UpdateMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequestBody", string(data)}, " ")
}
