package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyAllProjectRole
type AgencyAllProjectRole struct {

	// 权限ID。
	Id string `json:"id"`

	Links *LinksSelf `json:"links"`

	// 权限名。
	Name string `json:"name"`
}

func (o AgencyAllProjectRole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAllProjectRole struct{}"
	}

	return strings.Join([]string{"AgencyAllProjectRole", string(data)}, " ")
}
