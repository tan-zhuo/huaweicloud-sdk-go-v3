package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpClassificationItem struct {

	// IpItem的总数量
	Total *int32 `json:"total,omitempty"`

	// IpItem详细信息
	Items *[]IpItem `json:"items,omitempty"`
}

func (o IpClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpClassificationItem struct{}"
	}

	return strings.Join([]string{"IpClassificationItem", string(data)}, " ")
}
