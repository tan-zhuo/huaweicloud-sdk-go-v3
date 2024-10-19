package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateLifeCycleHookResponse Response Object
type UpdateLifeCycleHookResponse struct {

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	// 生命周期挂钩类型。INSTANCE_TERMINATING;INSTANCE_LAUNCHING
	LifecycleHookType *UpdateLifeCycleHookResponseLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	// 生命周期挂钩默认回调操作。ABANDON;CONTINUE
	DefaultResult *UpdateLifeCycleHookResponseDefaultResult `json:"default_result,omitempty"`

	// 生命周期挂钩超时时间，单位秒。
	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	// SMN服务中Topic的唯一的资源标识。
	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	// SMN服务中Topic的资源名称。
	NotificationTopicName *string `json:"notification_topic_name,omitempty"`

	// 自定义通知消息。
	NotificationMetadata *string `json:"notification_metadata,omitempty"`

	// 生命周期挂钩创建时间，遵循UTC时间。
	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLifeCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLifeCycleHookResponse struct{}"
	}

	return strings.Join([]string{"UpdateLifeCycleHookResponse", string(data)}, " ")
}

type UpdateLifeCycleHookResponseLifecycleHookType struct {
	value string
}

type UpdateLifeCycleHookResponseLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING UpdateLifeCycleHookResponseLifecycleHookType
	INSTANCE_LAUNCHING   UpdateLifeCycleHookResponseLifecycleHookType
}

func GetUpdateLifeCycleHookResponseLifecycleHookTypeEnum() UpdateLifeCycleHookResponseLifecycleHookTypeEnum {
	return UpdateLifeCycleHookResponseLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: UpdateLifeCycleHookResponseLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: UpdateLifeCycleHookResponseLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c UpdateLifeCycleHookResponseLifecycleHookType) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookResponseLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookResponseLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type UpdateLifeCycleHookResponseDefaultResult struct {
	value string
}

type UpdateLifeCycleHookResponseDefaultResultEnum struct {
	ABANDON  UpdateLifeCycleHookResponseDefaultResult
	CONTINUE UpdateLifeCycleHookResponseDefaultResult
}

func GetUpdateLifeCycleHookResponseDefaultResultEnum() UpdateLifeCycleHookResponseDefaultResultEnum {
	return UpdateLifeCycleHookResponseDefaultResultEnum{
		ABANDON: UpdateLifeCycleHookResponseDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: UpdateLifeCycleHookResponseDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c UpdateLifeCycleHookResponseDefaultResult) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookResponseDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookResponseDefaultResult) UnmarshalJSON(b []byte) error {
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
