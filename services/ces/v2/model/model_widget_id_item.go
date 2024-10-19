package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WidgetIdItem struct {

	// 视图id
	WidgetId *string `json:"widget_id,omitempty"`
}

func (o WidgetIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetIdItem struct{}"
	}

	return strings.Join([]string{"WidgetIdItem", string(data)}, " ")
}
