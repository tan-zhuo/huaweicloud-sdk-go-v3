package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrgInstancesResponse Response Object
type ListOrgInstancesResponse struct {
	Instances *PageInstancesVo `json:"instances,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOrgInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrgInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListOrgInstancesResponse", string(data)}, " ")
}
