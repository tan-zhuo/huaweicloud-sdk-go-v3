package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineSupportFeaturesEntity 实例支持的功能特性。
type ListEngineSupportFeaturesEntity struct {

	// 功能名称。
	Name *string `json:"name,omitempty"`

	Properties *ListEngineSupportFeaturesPropertiesEntity `json:"properties,omitempty"`
}

func (o ListEngineSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"ListEngineSupportFeaturesEntity", string(data)}, " ")
}
