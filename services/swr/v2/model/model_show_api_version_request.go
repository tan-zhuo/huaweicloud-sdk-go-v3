package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiVersionRequest Request Object
type ShowApiVersionRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowApiVersionRequestContentType `json:"Content-Type"`

	// API版本号。
	ApiVersion string `json:"api_version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}

type ShowApiVersionRequestContentType struct {
	value string
}

type ShowApiVersionRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowApiVersionRequestContentType
	APPLICATION_JSON             ShowApiVersionRequestContentType
}

func GetShowApiVersionRequestContentTypeEnum() ShowApiVersionRequestContentTypeEnum {
	return ShowApiVersionRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowApiVersionRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowApiVersionRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowApiVersionRequestContentType) Value() string {
	return c.value
}

func (c ShowApiVersionRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionRequestContentType) UnmarshalJSON(b []byte) error {
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
