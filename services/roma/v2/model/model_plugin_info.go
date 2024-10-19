package model

import (
	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type PluginInfo struct {

	// 插件编码。
	PluginId string `json:"plugin_id"`

	// 插件名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符。 > 中文字符必须为UTF-8或者unicode编码。
	PluginName string `json:"plugin_name"`

	// 插件类型 - cors：跨域资源共享 - set_resp_headers：HTTP响应头管理 - kafka_log：Kafka日志推送 - breaker：断路器 - rate_limit: 流量控制 - third_auth: 第三方认证
	PluginType PluginInfoPluginType `json:"plugin_type"`

	// 插件可见范围 - global：全局可见 - app：集成应用可见
	PluginScope PluginInfoPluginScope `json:"plugin_scope"`

	// 插件定义内容，支持json。参考提供的具体模型定义  CorsPluginContent：跨域资源共享 定义内容 SetRespHeadersContent：HTTP响应头管理 定义内容 KafkaLogContent：Kafka日志推送 定义内容 BreakerContent：断路器 定义内容 RateLimitContent 流量控制 定义内容 ThirdAuthContent: 第三方认证 定义内容
	PluginContent string `json:"plugin_content"`

	// 插件描述，255字符。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 归属集成应用编码，plugin_scope为app时生效
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// API归属的集成应用名称
	RomaAppName *string `json:"roma_app_name,omitempty"`
}

func (o PluginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginInfo struct{}"
	}

	return strings.Join([]string{"PluginInfo", string(data)}, " ")
}

type PluginInfoPluginType struct {
	value string
}

type PluginInfoPluginTypeEnum struct {
	CORS             PluginInfoPluginType
	SET_RESP_HEADERS PluginInfoPluginType
	KAFKA_LOG        PluginInfoPluginType
	BREAKER          PluginInfoPluginType
	RATE_LIMIT       PluginInfoPluginType
	THIRD_AUTH       PluginInfoPluginType
}

func GetPluginInfoPluginTypeEnum() PluginInfoPluginTypeEnum {
	return PluginInfoPluginTypeEnum{
		CORS: PluginInfoPluginType{
			value: "cors",
		},
		SET_RESP_HEADERS: PluginInfoPluginType{
			value: "set_resp_headers",
		},
		KAFKA_LOG: PluginInfoPluginType{
			value: "kafka_log",
		},
		BREAKER: PluginInfoPluginType{
			value: "breaker",
		},
		RATE_LIMIT: PluginInfoPluginType{
			value: "rate_limit",
		},
		THIRD_AUTH: PluginInfoPluginType{
			value: "third_auth",
		},
	}
}

func (c PluginInfoPluginType) Value() string {
	return c.value
}

func (c PluginInfoPluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginInfoPluginType) UnmarshalJSON(b []byte) error {
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

type PluginInfoPluginScope struct {
	value string
}

type PluginInfoPluginScopeEnum struct {
	GLOBAL PluginInfoPluginScope
	APP    PluginInfoPluginScope
}

func GetPluginInfoPluginScopeEnum() PluginInfoPluginScopeEnum {
	return PluginInfoPluginScopeEnum{
		GLOBAL: PluginInfoPluginScope{
			value: "global",
		},
		APP: PluginInfoPluginScope{
			value: "app",
		},
	}
}

func (c PluginInfoPluginScope) Value() string {
	return c.value
}

func (c PluginInfoPluginScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginInfoPluginScope) UnmarshalJSON(b []byte) error {
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
