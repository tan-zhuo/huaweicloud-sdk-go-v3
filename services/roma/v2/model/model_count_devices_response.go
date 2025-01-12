package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountDevicesResponse Response Object
type CountDevicesResponse struct {
	ProductTemplates *ProductTemplatesCalculation `json:"product_templates,omitempty"`

	Products *ProductsCalculation `json:"products,omitempty"`

	Devices        *DevicesCalculation `json:"devices,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CountDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDevicesResponse struct{}"
	}

	return strings.Join([]string{"CountDevicesResponse", string(data)}, " ")
}
