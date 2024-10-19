package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MonitorSystemRequestBodySpec 创建或更新监控系统配置。
type MonitorSystemRequestBodySpec struct {

	// 采集方式。
	Type MonitorSystemRequestBodySpecType `json:"type"`

	// 探针注入方式。
	Instrumentation MonitorSystemRequestBodySpecInstrumentation `json:"instrumentation"`

	// apm2访问密钥Key。
	AccessKey *string `json:"access_key,omitempty"`

	// apm2访问密钥value。
	AccessValue *string `json:"access_value,omitempty"`

	// apm2 opentelemetry接入token。
	AccessToken *string `json:"access_token,omitempty"`

	// apm2应用。
	ApmApplication string `json:"apm_application"`

	// 增强型探针/opentelemetry探针版本。
	Version string `json:"version"`

	// 探针镜像更新策略。
	ImagePullPolicy MonitorSystemRequestBodySpecImagePullPolicy `json:"image_pull_policy"`
}

func (o MonitorSystemRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorSystemRequestBodySpec struct{}"
	}

	return strings.Join([]string{"MonitorSystemRequestBodySpec", string(data)}, " ")
}

type MonitorSystemRequestBodySpecType struct {
	value string
}

type MonitorSystemRequestBodySpecTypeEnum struct {
	OPENTELEMETRY MonitorSystemRequestBodySpecType
	APM2          MonitorSystemRequestBodySpecType
}

func GetMonitorSystemRequestBodySpecTypeEnum() MonitorSystemRequestBodySpecTypeEnum {
	return MonitorSystemRequestBodySpecTypeEnum{
		OPENTELEMETRY: MonitorSystemRequestBodySpecType{
			value: "opentelemetry",
		},
		APM2: MonitorSystemRequestBodySpecType{
			value: "apm2",
		},
	}
}

func (c MonitorSystemRequestBodySpecType) Value() string {
	return c.value
}

func (c MonitorSystemRequestBodySpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorSystemRequestBodySpecType) UnmarshalJSON(b []byte) error {
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

type MonitorSystemRequestBodySpecInstrumentation struct {
	value string
}

type MonitorSystemRequestBodySpecInstrumentationEnum struct {
	MANUAL    MonitorSystemRequestBodySpecInstrumentation
	AUTOMATIC MonitorSystemRequestBodySpecInstrumentation
}

func GetMonitorSystemRequestBodySpecInstrumentationEnum() MonitorSystemRequestBodySpecInstrumentationEnum {
	return MonitorSystemRequestBodySpecInstrumentationEnum{
		MANUAL: MonitorSystemRequestBodySpecInstrumentation{
			value: "manual",
		},
		AUTOMATIC: MonitorSystemRequestBodySpecInstrumentation{
			value: "automatic",
		},
	}
}

func (c MonitorSystemRequestBodySpecInstrumentation) Value() string {
	return c.value
}

func (c MonitorSystemRequestBodySpecInstrumentation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorSystemRequestBodySpecInstrumentation) UnmarshalJSON(b []byte) error {
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

type MonitorSystemRequestBodySpecImagePullPolicy struct {
	value string
}

type MonitorSystemRequestBodySpecImagePullPolicyEnum struct {
	ALWAYS         MonitorSystemRequestBodySpecImagePullPolicy
	IF_NOT_PRESENT MonitorSystemRequestBodySpecImagePullPolicy
}

func GetMonitorSystemRequestBodySpecImagePullPolicyEnum() MonitorSystemRequestBodySpecImagePullPolicyEnum {
	return MonitorSystemRequestBodySpecImagePullPolicyEnum{
		ALWAYS: MonitorSystemRequestBodySpecImagePullPolicy{
			value: "Always",
		},
		IF_NOT_PRESENT: MonitorSystemRequestBodySpecImagePullPolicy{
			value: "IfNotPresent",
		},
	}
}

func (c MonitorSystemRequestBodySpecImagePullPolicy) Value() string {
	return c.value
}

func (c MonitorSystemRequestBodySpecImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorSystemRequestBodySpecImagePullPolicy) UnmarshalJSON(b []byte) error {
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
