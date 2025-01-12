package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisableAction 停用企业项目操作
type DisableAction struct {

	// 停用操作
	Action DisableActionAction `json:"action"`
}

func (o DisableAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAction struct{}"
	}

	return strings.Join([]string{"DisableAction", string(data)}, " ")
}

type DisableActionAction struct {
	value string
}

type DisableActionActionEnum struct {
	DISABLE DisableActionAction
}

func GetDisableActionActionEnum() DisableActionActionEnum {
	return DisableActionActionEnum{
		DISABLE: DisableActionAction{
			value: "disable",
		},
	}
}

func (c DisableActionAction) Value() string {
	return c.value
}

func (c DisableActionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisableActionAction) UnmarshalJSON(b []byte) error {
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
