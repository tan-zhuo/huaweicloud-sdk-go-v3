package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgentInvocationsRequest Request Object
type ListAgentInvocationsRequest struct {

	// 主机id
	InstanceId *string `json:"instance_id,omitempty"`

	// 主机类型，ECS弹性云服务器，BMS裸金属服务器
	InstanceType *ListAgentInvocationsRequestInstanceType `json:"instance_type,omitempty"`

	// 任务id
	InvocationId *string `json:"invocation_id,omitempty"`

	// 任务类型，INSTALL 安装，UPDATE 升级，ROLLBACK 回退，RETRY 重试
	InvocationType *ListAgentInvocationsRequestInvocationType `json:"invocation_type,omitempty"`

	// 任务对象，目前仅支持telescope
	InvocationTarget *ListAgentInvocationsRequestInvocationTarget `json:"invocation_target,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentInvocationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInvocationsRequest struct{}"
	}

	return strings.Join([]string{"ListAgentInvocationsRequest", string(data)}, " ")
}

type ListAgentInvocationsRequestInstanceType struct {
	value string
}

type ListAgentInvocationsRequestInstanceTypeEnum struct {
	ECS ListAgentInvocationsRequestInstanceType
	BMS ListAgentInvocationsRequestInstanceType
}

func GetListAgentInvocationsRequestInstanceTypeEnum() ListAgentInvocationsRequestInstanceTypeEnum {
	return ListAgentInvocationsRequestInstanceTypeEnum{
		ECS: ListAgentInvocationsRequestInstanceType{
			value: "ECS",
		},
		BMS: ListAgentInvocationsRequestInstanceType{
			value: "BMS",
		},
	}
}

func (c ListAgentInvocationsRequestInstanceType) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInstanceType) UnmarshalJSON(b []byte) error {
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

type ListAgentInvocationsRequestInvocationType struct {
	value string
}

type ListAgentInvocationsRequestInvocationTypeEnum struct {
	INSTALL  ListAgentInvocationsRequestInvocationType
	UPDATE   ListAgentInvocationsRequestInvocationType
	ROLLBACK ListAgentInvocationsRequestInvocationType
	RETRY    ListAgentInvocationsRequestInvocationType
}

func GetListAgentInvocationsRequestInvocationTypeEnum() ListAgentInvocationsRequestInvocationTypeEnum {
	return ListAgentInvocationsRequestInvocationTypeEnum{
		INSTALL: ListAgentInvocationsRequestInvocationType{
			value: "INSTALL",
		},
		UPDATE: ListAgentInvocationsRequestInvocationType{
			value: "UPDATE",
		},
		ROLLBACK: ListAgentInvocationsRequestInvocationType{
			value: "ROLLBACK",
		},
		RETRY: ListAgentInvocationsRequestInvocationType{
			value: "RETRY",
		},
	}
}

func (c ListAgentInvocationsRequestInvocationType) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInvocationType) UnmarshalJSON(b []byte) error {
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

type ListAgentInvocationsRequestInvocationTarget struct {
	value string
}

type ListAgentInvocationsRequestInvocationTargetEnum struct {
	TELESCOPE ListAgentInvocationsRequestInvocationTarget
}

func GetListAgentInvocationsRequestInvocationTargetEnum() ListAgentInvocationsRequestInvocationTargetEnum {
	return ListAgentInvocationsRequestInvocationTargetEnum{
		TELESCOPE: ListAgentInvocationsRequestInvocationTarget{
			value: "telescope",
		},
	}
}

func (c ListAgentInvocationsRequestInvocationTarget) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInvocationTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInvocationTarget) UnmarshalJSON(b []byte) error {
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
