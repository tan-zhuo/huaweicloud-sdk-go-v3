package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeasureUnitRest struct {

	// 度量单位ID。 例如：10表示GB。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 度量单位的名称，根据查询的语言类型返回结果。 例如：GB。
	MeasureName *string `json:"measure_name,omitempty"`

	// 度量单位名称的英文缩写。 例如：度量单位名称“GB”的英文缩写为“G”。
	Abbreviation *string `json:"abbreviation,omitempty"`

	// 度量类型。 1：货币2：时长3：流量4：数量6：长度7：容量9：行数10：周期11：频率12：个数15：容量*时长16：带宽速率17：容量时长18：查询速率19：带宽速率（1000进制）20：性能测试用量21：面积22：视频23：吞吐量24：分辨率26：通用资源包抵扣单位27：核数*时长28：内存*时长29：IOPS*时长30：吞吐量*时长31：个/时长40：流量（1000进制）41：1K Tokens82：帧98：缓存带宽x时长
	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o MeasureUnitRest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeasureUnitRest struct{}"
	}

	return strings.Join([]string{"MeasureUnitRest", string(data)}, " ")
}
