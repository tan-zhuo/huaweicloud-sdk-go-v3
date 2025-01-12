package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyTargetParamsReq 修改数据库参数请求体
type ModifyTargetParamsReq struct {

	// 参数分组
	Group ModifyTargetParamsReqGroup `json:"group"`

	// 修改的参数信息
	Params []ParamsReqBean `json:"params"`
}

func (o ModifyTargetParamsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTargetParamsReq struct{}"
	}

	return strings.Join([]string{"ModifyTargetParamsReq", string(data)}, " ")
}

type ModifyTargetParamsReqGroup struct {
	value string
}

type ModifyTargetParamsReqGroupEnum struct {
	COMMON      ModifyTargetParamsReqGroup
	PERFORMANCE ModifyTargetParamsReqGroup
}

func GetModifyTargetParamsReqGroupEnum() ModifyTargetParamsReqGroupEnum {
	return ModifyTargetParamsReqGroupEnum{
		COMMON: ModifyTargetParamsReqGroup{
			value: "common",
		},
		PERFORMANCE: ModifyTargetParamsReqGroup{
			value: "performance",
		},
	}
}

func (c ModifyTargetParamsReqGroup) Value() string {
	return c.value
}

func (c ModifyTargetParamsReqGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyTargetParamsReqGroup) UnmarshalJSON(b []byte) error {
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
