package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateWidgetsResponse Response Object
type BatchUpdateWidgetsResponse struct {

	// 更新结果列表
	Widgets        *[]BatchUpdateWidgetInfo `json:"widgets,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchUpdateWidgetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateWidgetsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateWidgetsResponse", string(data)}, " ")
}
