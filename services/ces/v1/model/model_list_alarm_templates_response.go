package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplatesResponse Response Object
type ListAlarmTemplatesResponse struct {

	// 自定义告警模板详细信息。
	AlarmTemplates *[]AlarmTemplate `json:"alarm_templates,omitempty"`

	MetaData       *MetaData `json:"meta_data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAlarmTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesResponse", string(data)}, " ")
}
