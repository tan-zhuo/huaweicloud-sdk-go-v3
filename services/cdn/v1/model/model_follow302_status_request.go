package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Follow302StatusRequest This is a auto create Body Object
type Follow302StatusRequest struct {

	// follow302状态，off：关闭，on：开启。
	Follow302Status Follow302StatusRequestFollow302Status `json:"follow302_status"`
}

func (o Follow302StatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Follow302StatusRequest struct{}"
	}

	return strings.Join([]string{"Follow302StatusRequest", string(data)}, " ")
}

type Follow302StatusRequestFollow302Status struct {
	value string
}

type Follow302StatusRequestFollow302StatusEnum struct {
	OFF Follow302StatusRequestFollow302Status
	ON  Follow302StatusRequestFollow302Status
}

func GetFollow302StatusRequestFollow302StatusEnum() Follow302StatusRequestFollow302StatusEnum {
	return Follow302StatusRequestFollow302StatusEnum{
		OFF: Follow302StatusRequestFollow302Status{
			value: "off",
		},
		ON: Follow302StatusRequestFollow302Status{
			value: "on",
		},
	}
}

func (c Follow302StatusRequestFollow302Status) Value() string {
	return c.value
}

func (c Follow302StatusRequestFollow302Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Follow302StatusRequestFollow302Status) UnmarshalJSON(b []byte) error {
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
