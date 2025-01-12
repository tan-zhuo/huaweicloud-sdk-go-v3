package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateFirewallResp
type UpdateFirewallResp struct {

	// 网络ACL ID
	Id string `json:"id"`

	// 网络ACL状态。
	Status UpdateFirewallRespStatus `json:"status"`
}

func (o UpdateFirewallResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallResp struct{}"
	}

	return strings.Join([]string{"UpdateFirewallResp", string(data)}, " ")
}

type UpdateFirewallRespStatus struct {
	value string
}

type UpdateFirewallRespStatusEnum struct {
	INACTIVE UpdateFirewallRespStatus
}

func GetUpdateFirewallRespStatusEnum() UpdateFirewallRespStatusEnum {
	return UpdateFirewallRespStatusEnum{
		INACTIVE: UpdateFirewallRespStatus{
			value: "INACTIVE",
		},
	}
}

func (c UpdateFirewallRespStatus) Value() string {
	return c.value
}

func (c UpdateFirewallRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFirewallRespStatus) UnmarshalJSON(b []byte) error {
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
