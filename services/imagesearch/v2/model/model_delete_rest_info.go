package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRestInfo 删除数据的相关信息。
type DeleteRestInfo struct {

	// 删除数据列表。
	Items *[]DeleteRestInfoItems `json:"items,omitempty"`

	DeleteInfo *DeleteInfo `json:"delete_info,omitempty"`
}

func (o DeleteRestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRestInfo struct{}"
	}

	return strings.Join([]string{"DeleteRestInfo", string(data)}, " ")
}
