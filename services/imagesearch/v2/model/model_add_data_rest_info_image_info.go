package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDataRestInfoImageInfo 添加图像数据的相关信息，不同服务类型返回信息不同，具体可参见服务类型说明。
type AddDataRestInfoImageInfo struct {

	// 添加的主体列表。
	Objects *[]AddDataRestInfoImageInfoObjects `json:"objects,omitempty"`
}

func (o AddDataRestInfoImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataRestInfoImageInfo struct{}"
	}

	return strings.Join([]string{"AddDataRestInfoImageInfo", string(data)}, " ")
}
