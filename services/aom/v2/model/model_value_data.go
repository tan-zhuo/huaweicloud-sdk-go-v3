package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValueData 指标具体数值。
type ValueData struct {

	// 指标名称。长度1~255。
	MetricName string `json:"metric_name"`

	// 数据的类型。取值范围只能是\"int\"或\"float\"。
	Type *ValueDataType `json:"type,omitempty"`

	// 数据的单位。长度不超过32个字符。
	Unit *string `json:"unit,omitempty"`

	// 指标数据的值。取值范围有效的数值类型。
	Value float64 `json:"value"`
}

func (o ValueData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueData struct{}"
	}

	return strings.Join([]string{"ValueData", string(data)}, " ")
}

type ValueDataType struct {
	value string
}

type ValueDataTypeEnum struct {
	INT   ValueDataType
	FLOAT ValueDataType
}

func GetValueDataTypeEnum() ValueDataTypeEnum {
	return ValueDataTypeEnum{
		INT: ValueDataType{
			value: "int",
		},
		FLOAT: ValueDataType{
			value: "float",
		},
	}
}

func (c ValueDataType) Value() string {
	return c.value
}

func (c ValueDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValueDataType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
