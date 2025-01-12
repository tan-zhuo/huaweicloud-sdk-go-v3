package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartGraphRequest Request Object
type StartGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图actionId
	ActionId StartGraphRequestActionId `json:"action_id"`

	Body *StartGraphReq `json:"body,omitempty"`
}

func (o StartGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartGraphRequest struct{}"
	}

	return strings.Join([]string{"StartGraphRequest", string(data)}, " ")
}

type StartGraphRequestActionId struct {
	value string
}

type StartGraphRequestActionIdEnum struct {
	START StartGraphRequestActionId
}

func GetStartGraphRequestActionIdEnum() StartGraphRequestActionIdEnum {
	return StartGraphRequestActionIdEnum{
		START: StartGraphRequestActionId{
			value: "start",
		},
	}
}

func (c StartGraphRequestActionId) Value() string {
	return c.value
}

func (c StartGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
