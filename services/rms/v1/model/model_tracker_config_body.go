package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrackerConfigBody TrackerConfig对象
type TrackerConfigBody struct {
	Channel *ChannelConfigBody `json:"channel"`

	Selector *SelectorConfigBody `json:"selector"`

	// IAM委托名称
	AgencyName string `json:"agency_name"`
}

func (o TrackerConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerConfigBody", string(data)}, " ")
}
