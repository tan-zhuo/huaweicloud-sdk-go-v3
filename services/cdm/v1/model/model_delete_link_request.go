package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLinkRequest Request Object
type DeleteLinkRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 需要删除的连接名
	LinkName string `json:"link_name"`
}

func (o DeleteLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLinkRequest struct{}"
	}

	return strings.Join([]string{"DeleteLinkRequest", string(data)}, " ")
}
