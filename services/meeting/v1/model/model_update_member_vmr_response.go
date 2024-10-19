package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberVmrResponse Response Object
type UpdateMemberVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMemberVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberVmrResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberVmrResponse", string(data)}, " ")
}
