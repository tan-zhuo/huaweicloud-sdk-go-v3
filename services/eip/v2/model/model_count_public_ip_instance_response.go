package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountPublicIpInstanceResponse Response Object
type CountPublicIpInstanceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountPublicIpInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPublicIpInstanceResponse struct{}"
	}

	return strings.Join([]string{"CountPublicIpInstanceResponse", string(data)}, " ")
}
