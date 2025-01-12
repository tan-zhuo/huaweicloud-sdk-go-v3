package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcTagResponse Response Object
type DeleteVpcTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcTagResponse", string(data)}, " ")
}
