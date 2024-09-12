package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveRoomInteractionRuleInfo 直播使用互动规则配置信息
type LiveRoomInteractionRuleInfo struct {

	// **参数解释**： 规则索引。用于触发规则时索引具体规则。 **约束限制**： 无需用户填写。 **取值范围**： 字符0-64位 **默认取值**： 不涉及。
	RuleIndex *string `json:"rule_index,omitempty"`

	// **参数解释**： 规则名称。 **约束限制**： 不涉及。 **取值范围**： 字符0-256位 **默认取值**： 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**： 是否启用。 **约束限制**： 不涉及。 **取值范围**： * true：启用 * fasle：不启用  **默认取值**： true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 规则匹配直播事件类型。接口的取值范围[0,100]，实际业务取值如下所示： * 1：弹幕事件  * 2：用户入场事件  * 3：用户点赞事件 * 4：用户送礼事件  * 10: 预置话术事件  请以实际业务取值为准。 > * 1,2,3,4：与LiveEventReport中的event.type对应。 > * 10：匹配预置剧本  **约束限制**： 不涉及。 **默认取值**： 不涉及
	EventType *int32 `json:"event_type,omitempty"`

	HitCondition *HitCondition `json:"hit_condition,omitempty"`

	Trigger *TriggerProcess `json:"trigger,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// **参数解释**： 互动规则库ID。从互动库导入时携带互动库ID。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-64位。 **默认取值**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 互动规则库名称。从互动库导入时携带互动库名称。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 规则启用场景。启用部分场景时填写。如果不填写则适用所有场景。 **约束限制**： 不涉及。 **取值范围**： 对应直播间场景列表，最大100个场景，这边也对应100个。 条目填写场景名称，字符长度1-256。 **默认取值**： 不涉及。
	EnabledScenes *[]string `json:"enabled_scenes,omitempty"`
}

func (o LiveRoomInteractionRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveRoomInteractionRuleInfo struct{}"
	}

	return strings.Join([]string{"LiveRoomInteractionRuleInfo", string(data)}, " ")
}
