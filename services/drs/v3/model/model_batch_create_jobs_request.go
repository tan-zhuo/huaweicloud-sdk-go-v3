package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateJobsRequest Request Object
type BatchCreateJobsRequest struct {

	// 请求语言类型
	XLanguage *BatchCreateJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchCreateJobReq `json:"body,omitempty"`
}

func (o BatchCreateJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsRequest", string(data)}, " ")
}

type BatchCreateJobsRequestXLanguage struct {
	value string
}

type BatchCreateJobsRequestXLanguageEnum struct {
	EN_US BatchCreateJobsRequestXLanguage
	ZH_CN BatchCreateJobsRequestXLanguage
}

func GetBatchCreateJobsRequestXLanguageEnum() BatchCreateJobsRequestXLanguageEnum {
	return BatchCreateJobsRequestXLanguageEnum{
		EN_US: BatchCreateJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCreateJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCreateJobsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchCreateJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
