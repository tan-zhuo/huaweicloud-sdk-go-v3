package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConfigurationsDiffRequest Request Object
type ListConfigurationsDiffRequest struct {

	// 语言。
	XLanguage *ListConfigurationsDiffRequestXLanguage `json:"X-Language,omitempty"`

	Body *ParamGroupDiffRequestBody `json:"body,omitempty"`
}

func (o ListConfigurationsDiffRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsDiffRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsDiffRequest", string(data)}, " ")
}

type ListConfigurationsDiffRequestXLanguage struct {
	value string
}

type ListConfigurationsDiffRequestXLanguageEnum struct {
	ZH_CN ListConfigurationsDiffRequestXLanguage
	EN_US ListConfigurationsDiffRequestXLanguage
}

func GetListConfigurationsDiffRequestXLanguageEnum() ListConfigurationsDiffRequestXLanguageEnum {
	return ListConfigurationsDiffRequestXLanguageEnum{
		ZH_CN: ListConfigurationsDiffRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListConfigurationsDiffRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListConfigurationsDiffRequestXLanguage) Value() string {
	return c.value
}

func (c ListConfigurationsDiffRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConfigurationsDiffRequestXLanguage) UnmarshalJSON(b []byte) error {
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
