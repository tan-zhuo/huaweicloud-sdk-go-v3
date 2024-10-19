package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobKindObj API类型，固定值“Job”，该值不可修改。
type JobKindObj struct {
	value string
}

type JobKindObjEnum struct {
	JOB JobKindObj
}

func GetJobKindObjEnum() JobKindObjEnum {
	return JobKindObjEnum{
		JOB: JobKindObj{
			value: "Job",
		},
	}
}

func (c JobKindObj) Value() string {
	return c.value
}

func (c JobKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobKindObj) UnmarshalJSON(b []byte) error {
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
