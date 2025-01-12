package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyV2XEdgeDto V2XEdge模型资源返回对象
type ModifyV2XEdgeDto struct {

	// **参数说明**：名称。  **取值范围**：长度不超过128，只允许中文、字母、数字、以及_.-等字符的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：Edge描述。  **取值范围**：长度不超过255，只允许中文、字母、数字、以及_.-等字符的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：设备编码。仅用于一致性检查，不可修改。  **取值范围**：长度不超过64，只允许中文、字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：网络IP，例如：127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：ITS800,ATLAS 端口号。
	Port *int32 `json:"port,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	Location *Location `json:"location,omitempty"`

	// **参数说明**：摄像头ID列表。
	CameraIds *[]string `json:"camera_ids,omitempty"`

	// **参数说明**：雷达ID列表。
	RadarIds *[]string `json:"radar_ids,omitempty"`

	// **参数说明**：Edge关联的本地RSU列表。
	LocalRsus *[]string `json:"local_rsus,omitempty"`

	EdgeGeneralConfig *EdgeGeneralConfig `json:"edge_general_config,omitempty"`

	// **参数说明**：Edge高级配置（请谨慎修改，错误的配置将导致edge不可用），Json格式；如果想要删除整个edge_advance_config可以填写空Object（例如:{}）。如果想移除某个配置项，直接从配置数据中移除相应的属性即可。
	EdgeAdvanceConfig *interface{} `json:"edge_advance_config,omitempty"`
}

func (o ModifyV2XEdgeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyV2XEdgeDto struct{}"
	}

	return strings.Join([]string{"ModifyV2XEdgeDto", string(data)}, " ")
}
