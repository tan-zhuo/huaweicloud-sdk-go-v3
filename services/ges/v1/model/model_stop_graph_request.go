package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopGraphRequest Request Object
type StopGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图actionId
	ActionId StopGraphRequestActionId `json:"action_id"`
}

func (o StopGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopGraphRequest struct{}"
	}

	return strings.Join([]string{"StopGraphRequest", string(data)}, " ")
}

type StopGraphRequestActionId struct {
	value string
}

type StopGraphRequestActionIdEnum struct {
	STOP StopGraphRequestActionId
}

func GetStopGraphRequestActionIdEnum() StopGraphRequestActionIdEnum {
	return StopGraphRequestActionIdEnum{
		STOP: StopGraphRequestActionId{
			value: "stop",
		},
	}
}

func (c StopGraphRequestActionId) Value() string {
	return c.value
}

func (c StopGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
