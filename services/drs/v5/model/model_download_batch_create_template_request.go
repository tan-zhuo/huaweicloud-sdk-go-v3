package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadBatchCreateTemplateRequest Request Object
type DownloadBatchCreateTemplateRequest struct {

	// 请求语言类型。
	XLanguage *DownloadBatchCreateTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DownloadBatchCreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBatchCreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadBatchCreateTemplateRequest", string(data)}, " ")
}

type DownloadBatchCreateTemplateRequestXLanguage struct {
	value string
}

type DownloadBatchCreateTemplateRequestXLanguageEnum struct {
	EN_US DownloadBatchCreateTemplateRequestXLanguage
	ZH_CN DownloadBatchCreateTemplateRequestXLanguage
}

func GetDownloadBatchCreateTemplateRequestXLanguageEnum() DownloadBatchCreateTemplateRequestXLanguageEnum {
	return DownloadBatchCreateTemplateRequestXLanguageEnum{
		EN_US: DownloadBatchCreateTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DownloadBatchCreateTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DownloadBatchCreateTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadBatchCreateTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadBatchCreateTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
