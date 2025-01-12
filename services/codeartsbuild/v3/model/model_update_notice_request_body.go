package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNoticeRequestBody 更新消息接口返回体
type UpdateNoticeRequestBody struct {

	// 通知类型
	NoticeType string `json:"notice_type"`

	// 开启的通知的种类
	EnabledEventTypeNames []string `json:"enabled_event_type_names"`

	// 通知参数配置
	ParamConfig *string `json:"param_config,omitempty"`
}

func (o UpdateNoticeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoticeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNoticeRequestBody", string(data)}, " ")
}
