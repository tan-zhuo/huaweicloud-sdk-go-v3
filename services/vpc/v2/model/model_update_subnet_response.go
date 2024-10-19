package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetResponse Response Object
type UpdateSubnetResponse struct {
	Subnet         *SubnetResult `json:"subnet,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponse", string(data)}, " ")
}
