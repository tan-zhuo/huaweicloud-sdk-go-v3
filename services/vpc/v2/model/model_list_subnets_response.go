package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetsResponse Response Object
type ListSubnetsResponse struct {

	// subnet对象列表
	Subnets        *[]Subnet `json:"subnets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsResponse", string(data)}, " ")
}
