package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBacklogPolicy 新增数据流转积压策略请求结构体
type AddBacklogPolicy struct {

	// **参数说明**：数据流转积压策略名称。 **取值范围**：长度不超过256，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合。
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数说明**：用户自定义的数据流转积压策略描述。 **取值范围**：长度不超过256，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：数据积压大小。单位为B（字节），取值范围为0~1073741823的整数，默认为1073741823（即1GB）。当backlog_size为0时，表示不积压。若同时配置了backlog_size和backlog_time两个维度，则以最先达到阈值的维度为准。
	BacklogSize *int32 `json:"backlog_size,omitempty"`

	// **参数说明**：数据积压时间。单位为s（秒），取值范围为0~86399的整数，默认为86399（即1天）。当backlog_time为0时，表示不积压。若同时配置了backlog_size和backlog_time两个维度，则以最先达到阈值的维度为准。
	BacklogTime *int32 `json:"backlog_time,omitempty"`
}

func (o AddBacklogPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBacklogPolicy struct{}"
	}

	return strings.Join([]string{"AddBacklogPolicy", string(data)}, " ")
}
