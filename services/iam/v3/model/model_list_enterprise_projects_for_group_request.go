package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectsForGroupRequest Request Object
type ListEnterpriseProjectsForGroupRequest struct {

	// 待查询用户组ID。
	GroupId string `json:"group_id"`
}

func (o ListEnterpriseProjectsForGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectsForGroupRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectsForGroupRequest", string(data)}, " ")
}
