package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGaussMySqlProjectQuotasRequest Request Object
type ShowGaussMySqlProjectQuotasRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 功能说明：根据type过滤查询指定类型的配额。  取值范围：instance
	Type *ShowGaussMySqlProjectQuotasRequestType `json:"type,omitempty"`
}

func (o ShowGaussMySqlProjectQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProjectQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProjectQuotasRequest", string(data)}, " ")
}

type ShowGaussMySqlProjectQuotasRequestType struct {
	value string
}

type ShowGaussMySqlProjectQuotasRequestTypeEnum struct {
	INSTANCE ShowGaussMySqlProjectQuotasRequestType
}

func GetShowGaussMySqlProjectQuotasRequestTypeEnum() ShowGaussMySqlProjectQuotasRequestTypeEnum {
	return ShowGaussMySqlProjectQuotasRequestTypeEnum{
		INSTANCE: ShowGaussMySqlProjectQuotasRequestType{
			value: "instance",
		},
	}
}

func (c ShowGaussMySqlProjectQuotasRequestType) Value() string {
	return c.value
}

func (c ShowGaussMySqlProjectQuotasRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGaussMySqlProjectQuotasRequestType) UnmarshalJSON(b []byte) error {
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
