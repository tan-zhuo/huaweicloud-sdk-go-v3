package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OriginHostBody struct {

	// accelerate：选择加速域名作为回源host域名； customize：使用自定义的域名作为回源host域名；
	OriginHostType OriginHostBodyOriginHostType `json:"origin_host_type"`

	// 自定义回源域名，origin_host_type为 customize时传入该参数。
	CustomizeDomain *string `json:"customize_domain,omitempty"`
}

func (o OriginHostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginHostBody struct{}"
	}

	return strings.Join([]string{"OriginHostBody", string(data)}, " ")
}

type OriginHostBodyOriginHostType struct {
	value string
}

type OriginHostBodyOriginHostTypeEnum struct {
	ACCELERATE OriginHostBodyOriginHostType
	CUSTOMIZE  OriginHostBodyOriginHostType
}

func GetOriginHostBodyOriginHostTypeEnum() OriginHostBodyOriginHostTypeEnum {
	return OriginHostBodyOriginHostTypeEnum{
		ACCELERATE: OriginHostBodyOriginHostType{
			value: "accelerate",
		},
		CUSTOMIZE: OriginHostBodyOriginHostType{
			value: "customize",
		},
	}
}

func (c OriginHostBodyOriginHostType) Value() string {
	return c.value
}

func (c OriginHostBodyOriginHostType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OriginHostBodyOriginHostType) UnmarshalJSON(b []byte) error {
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
