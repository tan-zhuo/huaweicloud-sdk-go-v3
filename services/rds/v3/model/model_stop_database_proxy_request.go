package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StopDatabaseProxyRequest struct {
	XLanguage  *StopDatabaseProxyRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                             `json:"instance_id"`
}

func (o StopDatabaseProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopDatabaseProxyRequest struct{}"
	}

	return strings.Join([]string{"StopDatabaseProxyRequest", string(data)}, " ")
}

type StopDatabaseProxyRequestXLanguage struct {
	value string
}

type StopDatabaseProxyRequestXLanguageEnum struct {
	ZH_CN StopDatabaseProxyRequestXLanguage
	EN_US StopDatabaseProxyRequestXLanguage
}

func GetStopDatabaseProxyRequestXLanguageEnum() StopDatabaseProxyRequestXLanguageEnum {
	return StopDatabaseProxyRequestXLanguageEnum{
		ZH_CN: StopDatabaseProxyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopDatabaseProxyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopDatabaseProxyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StopDatabaseProxyRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}