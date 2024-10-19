package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportConnectionReq struct {
	Path *string `json:"path,omitempty"`

	// 连接参数
	Params *[]ConnectionParam `json:"params,omitempty"`

	SameNamePolicy *ImportConnectionReqSameNamePolicy `json:"sameNamePolicy,omitempty"`
}

func (o ImportConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportConnectionReq struct{}"
	}

	return strings.Join([]string{"ImportConnectionReq", string(data)}, " ")
}

type ImportConnectionReqSameNamePolicy struct {
	value string
}

type ImportConnectionReqSameNamePolicyEnum struct {
	SKIP      ImportConnectionReqSameNamePolicy
	OVERWRITE ImportConnectionReqSameNamePolicy
}

func GetImportConnectionReqSameNamePolicyEnum() ImportConnectionReqSameNamePolicyEnum {
	return ImportConnectionReqSameNamePolicyEnum{
		SKIP: ImportConnectionReqSameNamePolicy{
			value: "SKIP",
		},
		OVERWRITE: ImportConnectionReqSameNamePolicy{
			value: "OVERWRITE",
		},
	}
}

func (c ImportConnectionReqSameNamePolicy) Value() string {
	return c.value
}

func (c ImportConnectionReqSameNamePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportConnectionReqSameNamePolicy) UnmarshalJSON(b []byte) error {
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
