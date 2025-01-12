package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyTokenAuth
type AgencyTokenAuth struct {
	Identity *AgencyTokenIdentity `json:"identity"`

	Scope *AgencyTokenScope `json:"scope"`
}

func (o AgencyTokenAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenAuth struct{}"
	}

	return strings.Join([]string{"AgencyTokenAuth", string(data)}, " ")
}
