package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupsRequest Request Object
type ListGroupsRequest struct {

	// 存储用量单位
	Unit *ListGroupsRequestUnit `json:"unit,omitempty"`

	// 存储类型，有资产存储(取值:AssetStorage)、设备存储(取值:DeviceStorage)两种类型
	Type *string `json:"type,omitempty"`

	// 存储组 ID
	GroupId *string `json:"group_id,omitempty"`

	// 存储组名称
	Name *string `json:"name,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupsRequest", string(data)}, " ")
}

type ListGroupsRequestUnit struct {
	value string
}

type ListGroupsRequestUnitEnum struct {
	MB ListGroupsRequestUnit
}

func GetListGroupsRequestUnitEnum() ListGroupsRequestUnitEnum {
	return ListGroupsRequestUnitEnum{
		MB: ListGroupsRequestUnit{
			value: "MB",
		},
	}
}

func (c ListGroupsRequestUnit) Value() string {
	return c.value
}

func (c ListGroupsRequestUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupsRequestUnit) UnmarshalJSON(b []byte) error {
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
