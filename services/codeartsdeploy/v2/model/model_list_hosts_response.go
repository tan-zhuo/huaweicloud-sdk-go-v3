package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsResponse Response Object
type ListHostsResponse struct {

	// 主机数量
	Total *int32 `json:"total,omitempty"`

	// 主机集群名称
	GroupName *string `json:"group_name,omitempty"`

	// 主机列表信息
	Hosts          *[]DeploymentHostDetail `json:"hosts,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
