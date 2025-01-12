package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchListStructProcessRequest Request Object
type BatchListStructProcessRequest struct {

	// 请求语言类型
	XLanguage *BatchListStructProcessRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryJobReq `json:"body,omitempty"`
}

func (o BatchListStructProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructProcessRequest struct{}"
	}

	return strings.Join([]string{"BatchListStructProcessRequest", string(data)}, " ")
}

type BatchListStructProcessRequestXLanguage struct {
	value string
}

type BatchListStructProcessRequestXLanguageEnum struct {
	EN_US BatchListStructProcessRequestXLanguage
	ZH_CN BatchListStructProcessRequestXLanguage
}

func GetBatchListStructProcessRequestXLanguageEnum() BatchListStructProcessRequestXLanguageEnum {
	return BatchListStructProcessRequestXLanguageEnum{
		EN_US: BatchListStructProcessRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListStructProcessRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListStructProcessRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListStructProcessRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListStructProcessRequestXLanguage) UnmarshalJSON(b []byte) error {
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
