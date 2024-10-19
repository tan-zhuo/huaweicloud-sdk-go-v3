package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnServersByProjectRequest Request Object
type ListVpnServersByProjectRequest struct {
}

func (o ListVpnServersByProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnServersByProjectRequest struct{}"
	}

	return strings.Join([]string{"ListVpnServersByProjectRequest", string(data)}, " ")
}
