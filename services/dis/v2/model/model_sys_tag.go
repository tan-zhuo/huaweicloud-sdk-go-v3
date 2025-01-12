package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SysTag struct {

	// 键。  - 不能为空。  - 值必须为_sys_enterprise_project_id。
	Key *SysTagKey `json:"key,omitempty"`

	// 值。  - 对应的是企业项目ID，需要在企业管理页面获取。  - 36位UUID。
	Value *string `json:"value,omitempty"`
}

func (o SysTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTag struct{}"
	}

	return strings.Join([]string{"SysTag", string(data)}, " ")
}

type SysTagKey struct {
	value string
}

type SysTagKeyEnum struct {
	SYS_ENTERPRISE_PROJECT_ID SysTagKey
}

func GetSysTagKeyEnum() SysTagKeyEnum {
	return SysTagKeyEnum{
		SYS_ENTERPRISE_PROJECT_ID: SysTagKey{
			value: "_sys_enterprise_project_id",
		},
	}
}

func (c SysTagKey) Value() string {
	return c.value
}

func (c SysTagKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SysTagKey) UnmarshalJSON(b []byte) error {
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
