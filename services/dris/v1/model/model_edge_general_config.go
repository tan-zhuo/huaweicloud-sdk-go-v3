package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EdgeGeneralConfig struct {

	// **参数说明**：AVP场景。
	AvpEnabled *bool `json:"avp_enabled,omitempty"`

	// **参数说明**：RSM上报：默认不上报。
	RsmEnabled *bool `json:"rsm_enabled,omitempty"`

	// **参数说明**：时延补偿：是否启动Edge时延补偿功能。
	TimeCompensate *bool `json:"time_compensate,omitempty"`

	// **参数说明**：RSI事件定位功能。
	RsiPositioningEnabled *bool `json:"rsi_positioning_enabled,omitempty"`

	// **参数说明**：应用日志级别，可选项：on/off，默认关闭。
	LogLevel *string `json:"log_level,omitempty"`

	// **参数说明**：道路检测长度，单位：米。
	RoadDetectionLength *float64 `json:"road_detection_length,omitempty"`

	// **参数说明**：匝道检测长度，单位：米。
	RampDetectionLength *float64 `json:"ramp_detection_length,omitempty"`

	// **参数说明**：1400接口用户名。  **取值范围**：长度不小于8，不大于32，只允许数字字母下划线组合，且不能以数字下划线开头，不能有中文和特殊字符，gat1400用户名不能与gat1400密码相同。
	Gat1400Username *string `json:"gat1400_username,omitempty"`

	// **参数说明**：1400接口密码。  **取值范围**：   1.密码必须包含如下至少两种字符的组合:     －一个小写字母；     －一个大写字母；     －一个数字；     －一个特殊字符：`~!@#$%^&*()-_=+\\|[{}];:'\",<.>/?  和空格   2.gat1400密码不能与gat1400用户名相同
	Gat1400Password *string `json:"gat1400_password,omitempty"`

	// **参数说明**：ITS800鉴权用的用户名和密码。
	UserName *string `json:"user_name,omitempty"`

	// **参数说明**：密码。
	Password *string `json:"password,omitempty"`
}

func (o EdgeGeneralConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGeneralConfig struct{}"
	}

	return strings.Join([]string{"EdgeGeneralConfig", string(data)}, " ")
}
