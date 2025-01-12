package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MatchOption struct {

	// 取值为“instance_name”或“instance_id”，分别表示按实例名称或按实例ID匹配查询。
	Key MatchOptionKey `json:"key"`

	// 待匹配的实例名称或实例ID。
	Value string `json:"value"`
}

func (o MatchOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchOption struct{}"
	}

	return strings.Join([]string{"MatchOption", string(data)}, " ")
}

type MatchOptionKey struct {
	value string
}

type MatchOptionKeyEnum struct {
	INSTANCE_NAME MatchOptionKey
	INSTANCE_ID   MatchOptionKey
}

func GetMatchOptionKeyEnum() MatchOptionKeyEnum {
	return MatchOptionKeyEnum{
		INSTANCE_NAME: MatchOptionKey{
			value: "instance_name",
		},
		INSTANCE_ID: MatchOptionKey{
			value: "instance_id",
		},
	}
}

func (c MatchOptionKey) Value() string {
	return c.value
}

func (c MatchOptionKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MatchOptionKey) UnmarshalJSON(b []byte) error {
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
