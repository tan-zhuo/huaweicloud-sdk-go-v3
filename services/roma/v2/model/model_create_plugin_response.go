package model

import (
	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreatePluginResponse Response Object
type CreatePluginResponse struct {

	// 插件编码。
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证
	PluginType *CreatePluginResponsePluginType `json:"plugin_type,omitempty"`

	// 插件可见范围 - global：全局可见 - app：集成应用可见
	PluginScope *CreatePluginResponsePluginScope `json:"plugin_scope,omitempty"`

	// 插件定义内容，支持json。参考提供的具体模型定义  CorsPluginContent：跨域资源共享 定义内容 SetRespHeadersContent：HTTP响应头管理 定义内容 KafkaLogContent：Kafka日志推送 定义内容 BreakerContent：断路器 定义内容 RateLimitContent 流量控制 定义内容 ThirdAuthContent: 第三方认证 定义内容
	PluginContent *string `json:"plugin_content,omitempty"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 归属集成应用编码，plugin_scope为app时生效
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// API归属的集成应用名称
	RomaAppName    *string `json:"roma_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginResponse", string(data)}, " ")
}

type CreatePluginResponsePluginType struct {
	value string
}

type CreatePluginResponsePluginTypeEnum struct {
	CORS             CreatePluginResponsePluginType
	SET_RESP_HEADERS CreatePluginResponsePluginType
	KAFKA_LOG        CreatePluginResponsePluginType
	BREAKER          CreatePluginResponsePluginType
	RATE_LIMIT       CreatePluginResponsePluginType
	THIRD_AUTH       CreatePluginResponsePluginType
}

func GetCreatePluginResponsePluginTypeEnum() CreatePluginResponsePluginTypeEnum {
	return CreatePluginResponsePluginTypeEnum{
		CORS: CreatePluginResponsePluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: CreatePluginResponsePluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: CreatePluginResponsePluginType{
			value: "kafka_log",
		},
		BREAKER: CreatePluginResponsePluginType{
			value: "breaker",
		},
		RATE_LIMIT: CreatePluginResponsePluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: CreatePluginResponsePluginType{
			value: "third_auth",
		},
	}
}

func (c CreatePluginResponsePluginType) Value() string {
	return c.value
}

func (c CreatePluginResponsePluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponsePluginType) UnmarshalJSON(b []byte) error {
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

type CreatePluginResponsePluginScope struct {
	value string
}

type CreatePluginResponsePluginScopeEnum struct {
	GLOBAL CreatePluginResponsePluginScope
	APP    CreatePluginResponsePluginScope
}

func GetCreatePluginResponsePluginScopeEnum() CreatePluginResponsePluginScopeEnum {
	return CreatePluginResponsePluginScopeEnum{
		GLOBAL: CreatePluginResponsePluginScope{
			value: "global",
		},
		APP: CreatePluginResponsePluginScope{
			value: "app",
		},
	}
}

func (c CreatePluginResponsePluginScope) Value() string {
	return c.value
}

func (c CreatePluginResponsePluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePluginResponsePluginScope) UnmarshalJSON(b []byte) error {
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
