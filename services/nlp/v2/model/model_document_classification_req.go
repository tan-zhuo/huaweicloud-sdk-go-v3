package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DocumentClassificationReq
type DocumentClassificationReq struct {

	// 输入的文档，最大长度10000, 长度超过10000字符截取前10000个字符。
	Content string `json:"content"`

	// 预留字段，支持的文本语言类型，当前只支持zh（中文），默认zh。
	Lang *DocumentClassificationReqLang `json:"lang,omitempty"`
}

func (o DocumentClassificationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentClassificationReq struct{}"
	}

	return strings.Join([]string{"DocumentClassificationReq", string(data)}, " ")
}

type DocumentClassificationReqLang struct {
	value string
}

type DocumentClassificationReqLangEnum struct {
	ZH DocumentClassificationReqLang
}

func GetDocumentClassificationReqLangEnum() DocumentClassificationReqLangEnum {
	return DocumentClassificationReqLangEnum{
		ZH: DocumentClassificationReqLang{
			value: "zh",
		},
	}
}

func (c DocumentClassificationReqLang) Value() string {
	return c.value
}

func (c DocumentClassificationReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DocumentClassificationReqLang) UnmarshalJSON(b []byte) error {
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
