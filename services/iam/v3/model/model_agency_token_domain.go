package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyTokenDomain
type AgencyTokenDomain struct {

	// 委托方A的账号名称。
	Name string `json:"name"`

	// 委托方A的账号ID。
	Id string `json:"id"`
}

func (o AgencyTokenDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTokenDomain struct{}"
	}

	return strings.Join([]string{"AgencyTokenDomain", string(data)}, " ")
}
