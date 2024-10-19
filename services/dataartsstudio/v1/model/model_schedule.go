package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Schedule 调度配置
type Schedule struct {

	// 调度类型 - EXECUTE_ONCE: 作业立即运行，只运行一次。 - CRON: 作业按指定频率周期执行。 - EVENT:  根据事件触发执行。
	Type ScheduleType `json:"type"`

	Cron *Cron `json:"cron,omitempty"`

	Event *Event `json:"event,omitempty"`
}

func (o Schedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Schedule struct{}"
	}

	return strings.Join([]string{"Schedule", string(data)}, " ")
}

type ScheduleType struct {
	value string
}

type ScheduleTypeEnum struct {
	EXECUTE_ONCE ScheduleType
	CRON         ScheduleType
	EVENT        ScheduleType
}

func GetScheduleTypeEnum() ScheduleTypeEnum {
	return ScheduleTypeEnum{
		EXECUTE_ONCE: ScheduleType{
			value: "EXECUTE_ONCE",
		},
		CRON: ScheduleType{
			value: "CRON",
		},
		EVENT: ScheduleType{
			value: "EVENT",
		},
	}
}

func (c ScheduleType) Value() string {
	return c.value
}

func (c ScheduleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleType) UnmarshalJSON(b []byte) error {
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
