package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperateTasksBean struct {

	// 操作类型 - start (启动) - stop (停止)
	ActionId OperateTasksBeanActionId `json:"action_id"`

	// 需要启动或者停止的任务ID列表
	List *[]TaskBean `json:"list,omitempty"`
}

func (o OperateTasksBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateTasksBean struct{}"
	}

	return strings.Join([]string{"OperateTasksBean", string(data)}, " ")
}

type OperateTasksBeanActionId struct {
	value string
}

type OperateTasksBeanActionIdEnum struct {
	START OperateTasksBeanActionId
	STOP  OperateTasksBeanActionId
}

func GetOperateTasksBeanActionIdEnum() OperateTasksBeanActionIdEnum {
	return OperateTasksBeanActionIdEnum{
		START: OperateTasksBeanActionId{
			value: "start",
		},
		STOP: OperateTasksBeanActionId{
			value: "stop",
		},
	}
}

func (c OperateTasksBeanActionId) Value() string {
	return c.value
}

func (c OperateTasksBeanActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperateTasksBeanActionId) UnmarshalJSON(b []byte) error {
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
