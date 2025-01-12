package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateNamespaceRequest Request Object
type CreateNamespaceRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateNamespaceRequestContentType `json:"Content-Type"`

	Body *CreateNamespaceRequestBody `json:"body,omitempty"`
}

func (o CreateNamespaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNamespaceRequest struct{}"
	}

	return strings.Join([]string{"CreateNamespaceRequest", string(data)}, " ")
}

type CreateNamespaceRequestContentType struct {
	value string
}

type CreateNamespaceRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateNamespaceRequestContentType
	APPLICATION_JSON             CreateNamespaceRequestContentType
}

func GetCreateNamespaceRequestContentTypeEnum() CreateNamespaceRequestContentTypeEnum {
	return CreateNamespaceRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateNamespaceRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateNamespaceRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateNamespaceRequestContentType) Value() string {
	return c.value
}

func (c CreateNamespaceRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNamespaceRequestContentType) UnmarshalJSON(b []byte) error {
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
