package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowStoredQueryResponse Response Object
type ShowStoredQueryResponse struct {

	// ResourceQL ID
	Id *string `json:"id,omitempty"`

	// ResourceQL 名字
	Name *string `json:"name,omitempty"`

	// 自定义查询类型，枚举值为“account”和“aggregator”。若取值为“account”，表示单帐号的自定义查询语句；若取值为“aggregator”，表示聚合器的自定义查询语句。默认值为“account”。
	Type *ShowStoredQueryResponseType `json:"type,omitempty"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty"`

	// ResourceQL 表达式
	Expression *string `json:"expression,omitempty"`

	// ResourceQL 创建时间
	Created *string `json:"created,omitempty"`

	// ResourceQL 更新时间
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"ShowStoredQueryResponse", string(data)}, " ")
}

type ShowStoredQueryResponseType struct {
	value string
}

type ShowStoredQueryResponseTypeEnum struct {
	ACCOUNT    ShowStoredQueryResponseType
	AGGREGATOR ShowStoredQueryResponseType
}

func GetShowStoredQueryResponseTypeEnum() ShowStoredQueryResponseTypeEnum {
	return ShowStoredQueryResponseTypeEnum{
		ACCOUNT: ShowStoredQueryResponseType{
			value: "account",
		},
		AGGREGATOR: ShowStoredQueryResponseType{
			value: "aggregator",
		},
	}
}

func (c ShowStoredQueryResponseType) Value() string {
	return c.value
}

func (c ShowStoredQueryResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStoredQueryResponseType) UnmarshalJSON(b []byte) error {
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
