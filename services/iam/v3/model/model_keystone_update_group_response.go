package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateGroupResponse Response Object
type KeystoneUpdateGroupResponse struct {
	Group          *KeystoneGroupResultWithLinksSelf `json:"group,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o KeystoneUpdateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateGroupResponse", string(data)}, " ")
}
