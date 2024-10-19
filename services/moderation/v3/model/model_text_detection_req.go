package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextDetectionReq 文本内容审核请求体
type TextDetectionReq struct {

	// 事件类型。  可选值如下：  nickname: 昵称  title: 标题  article: 帖⼦  comment: 评论  barrage: 弹幕  search: 搜索栏  profile: 个⼈简介
	EventType *string `json:"event_type,omitempty"`

	// 文本审核场景，可选值如下： terrorism: 暴恐 porn: 色情 ban: 违禁 abuse: 辱骂 ad: 广告 当categories缺省或为空时，默认审核terrorism、porn、ban、abuse、ad。
	Categories *[]string `json:"categories,omitempty"`

	// 检测时使用的自定义黑名单词库列表。自定义黑词库的创建和使用请参见[配置定义黑名单词库](https://support.huaweicloud.com/api-moderation/moderation_03_0027.html#moderation_03_0027__section12400140132318)。
	GlossaryNames *[]string `json:"glossary_names,omitempty"`

	Data *TextDetectionDataReq `json:"data"`

	// 检测时使用的自定义白名单词库列表。自定义白词库的创建和使用请参见[配置定义白名单词库](https://support.huaweicloud.com/api-moderation/moderation_03_0027.html#moderation_03_0027__section178844141394)。
	WhiteGlossaryNames *[]string `json:"white_glossary_names,omitempty"`

	// 自定义审核策略名称，可在控制台配置;如果请求参数中传了biz_type则优先使用biz_type,如果用户没传biz_type则event_type必须传。
	BizType *string `json:"biz_type,omitempty"`
}

func (o TextDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionReq struct{}"
	}

	return strings.Join([]string{"TextDetectionReq", string(data)}, " ")
}
