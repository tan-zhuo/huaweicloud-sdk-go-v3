package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateFunctionTriggerRequestBody struct {

	// 触发器类型。  - TIMER: 定时触发器。 - APIG: APIGW触发器。 - CTS: 云审计触发器，需要先开通云审计服务。 - DDS: 文档数据库触发器，需要开启函数vpc。 - DMS: 分布式消息服务触发器，需要配置dms委托。 - DIS: 数据接入服务触发器，需要配置dis委托。 - LTS: 云审计日志服务触发器，需要配置lts委托。 - OBS: 对象存储服务触发器。 - KAFKA: 专享版本kafka触发器。
	TriggerTypeCode CreateFunctionTriggerRequestBodyTriggerTypeCode `json:"trigger_type_code"`

	// 触发器状态，取值为ACTIVE,DISABLED。
	TriggerStatus *CreateFunctionTriggerRequestBodyTriggerStatus `json:"trigger_status,omitempty"`

	// 消息代码。
	EventTypeCode *string `json:"event_type_code,omitempty"`

	EventData *TriggerEventDataRequestBody `json:"event_data"`
}

func (o CreateFunctionTriggerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionTriggerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFunctionTriggerRequestBody", string(data)}, " ")
}

type CreateFunctionTriggerRequestBodyTriggerTypeCode struct {
	value string
}

type CreateFunctionTriggerRequestBodyTriggerTypeCodeEnum struct {
	TIMER            CreateFunctionTriggerRequestBodyTriggerTypeCode
	APIG             CreateFunctionTriggerRequestBodyTriggerTypeCode
	CTS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	DDS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	DMS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	DIS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	LTS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	OBS              CreateFunctionTriggerRequestBodyTriggerTypeCode
	SMN              CreateFunctionTriggerRequestBodyTriggerTypeCode
	KAFKA            CreateFunctionTriggerRequestBodyTriggerTypeCode
	RABBITMQ         CreateFunctionTriggerRequestBodyTriggerTypeCode
	DEDICATEDGATEWAY CreateFunctionTriggerRequestBodyTriggerTypeCode
	OPENSOURCEKAFKA  CreateFunctionTriggerRequestBodyTriggerTypeCode
	APIC             CreateFunctionTriggerRequestBodyTriggerTypeCode
	GAUSSMONGO       CreateFunctionTriggerRequestBodyTriggerTypeCode
	EVENTGRID        CreateFunctionTriggerRequestBodyTriggerTypeCode
}

func GetCreateFunctionTriggerRequestBodyTriggerTypeCodeEnum() CreateFunctionTriggerRequestBodyTriggerTypeCodeEnum {
	return CreateFunctionTriggerRequestBodyTriggerTypeCodeEnum{
		TIMER: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "TIMER",
		},
		APIG: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "APIG",
		},
		CTS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "CTS",
		},
		DDS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "DDS",
		},
		DMS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "DMS",
		},
		DIS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "DIS",
		},
		LTS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "LTS",
		},
		OBS: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "OBS",
		},
		SMN: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: CreateFunctionTriggerRequestBodyTriggerTypeCode{
			value: "EVENTGRID",
		},
	}
}

func (c CreateFunctionTriggerRequestBodyTriggerTypeCode) Value() string {
	return c.value
}

func (c CreateFunctionTriggerRequestBodyTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerRequestBodyTriggerTypeCode) UnmarshalJSON(b []byte) error {
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

type CreateFunctionTriggerRequestBodyTriggerStatus struct {
	value string
}

type CreateFunctionTriggerRequestBodyTriggerStatusEnum struct {
	ACTIVE   CreateFunctionTriggerRequestBodyTriggerStatus
	DISABLED CreateFunctionTriggerRequestBodyTriggerStatus
}

func GetCreateFunctionTriggerRequestBodyTriggerStatusEnum() CreateFunctionTriggerRequestBodyTriggerStatusEnum {
	return CreateFunctionTriggerRequestBodyTriggerStatusEnum{
		ACTIVE: CreateFunctionTriggerRequestBodyTriggerStatus{
			value: "ACTIVE",
		},
		DISABLED: CreateFunctionTriggerRequestBodyTriggerStatus{
			value: "DISABLED",
		},
	}
}

func (c CreateFunctionTriggerRequestBodyTriggerStatus) Value() string {
	return c.value
}

func (c CreateFunctionTriggerRequestBodyTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerRequestBodyTriggerStatus) UnmarshalJSON(b []byte) error {
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
