package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchStartJobsRequest Request Object
type BatchStartJobsRequest struct {

	// 请求语言类型
	XLanguage *BatchStartJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchStartJobReq `json:"body,omitempty"`
}

func (o BatchStartJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchStartJobsRequest", string(data)}, " ")
}

type BatchStartJobsRequestXLanguage struct {
	value string
}

type BatchStartJobsRequestXLanguageEnum struct {
	EN_US BatchStartJobsRequestXLanguage
	ZH_CN BatchStartJobsRequestXLanguage
}

func GetBatchStartJobsRequestXLanguageEnum() BatchStartJobsRequestXLanguageEnum {
	return BatchStartJobsRequestXLanguageEnum{
		EN_US: BatchStartJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchStartJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchStartJobsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchStartJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStartJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
