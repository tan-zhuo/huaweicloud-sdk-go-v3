package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopsTagsReq 批量桌面删除标签
type BatchDeleteDesktopsTagsReq struct {

	// 桌面（桌面携带标签列表）列表
	Desktops []DesktopTagsInfo `json:"desktops"`
}

func (o BatchDeleteDesktopsTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopsTagsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopsTagsReq", string(data)}, " ")
}
