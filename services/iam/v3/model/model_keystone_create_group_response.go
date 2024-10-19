package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateGroupResponse Response Object
type KeystoneCreateGroupResponse struct {
	Group          *KeystoneGroupResultWithLinksSelf `json:"group,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o KeystoneCreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateGroupResponse", string(data)}, " ")
}
