package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConnectionResponse Response Object
type ShowConnectionResponse struct {
	Name *string `json:"name,omitempty"`

	Type *ShowConnectionResponseType `json:"type,omitempty"`

	Config *interface{} `json:"config,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionResponse", string(data)}, " ")
}

type ShowConnectionResponseType struct {
	value string
}

type ShowConnectionResponseTypeEnum struct {
	DWS         ShowConnectionResponseType
	DLI         ShowConnectionResponseType
	SPARK_SQL   ShowConnectionResponseType
	HIVE        ShowConnectionResponseType
	RDS         ShowConnectionResponseType
	CLOUD_TABLE ShowConnectionResponseType
}

func GetShowConnectionResponseTypeEnum() ShowConnectionResponseTypeEnum {
	return ShowConnectionResponseTypeEnum{
		DWS: ShowConnectionResponseType{
			value: "DWS",
		},
		DLI: ShowConnectionResponseType{
			value: "DLI",
		},
		SPARK_SQL: ShowConnectionResponseType{
			value: "SparkSQL",
		},
		HIVE: ShowConnectionResponseType{
			value: "Hive",
		},
		RDS: ShowConnectionResponseType{
			value: "RDS",
		},
		CLOUD_TABLE: ShowConnectionResponseType{
			value: "CloudTable",
		},
	}
}

func (c ShowConnectionResponseType) Value() string {
	return c.value
}

func (c ShowConnectionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConnectionResponseType) UnmarshalJSON(b []byte) error {
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
