package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApiVersionsRequest Request Object
type ListApiVersionsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListApiVersionsRequestContentType `json:"Content-Type"`
}

func (o ListApiVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionsRequest", string(data)}, " ")
}

type ListApiVersionsRequestContentType struct {
	value string
}

type ListApiVersionsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListApiVersionsRequestContentType
	APPLICATION_JSON             ListApiVersionsRequestContentType
}

func GetListApiVersionsRequestContentTypeEnum() ListApiVersionsRequestContentTypeEnum {
	return ListApiVersionsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListApiVersionsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListApiVersionsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListApiVersionsRequestContentType) Value() string {
	return c.value
}

func (c ListApiVersionsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApiVersionsRequestContentType) UnmarshalJSON(b []byte) error {
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
