package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsResponse Response Object
type ListPortsResponse struct {

	// port列表对象
	Ports          *[]Port `json:"ports,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsResponse struct{}"
	}

	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}
