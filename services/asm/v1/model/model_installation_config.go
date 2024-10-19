package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstallationConfig struct {
	Nodes *Selector `json:"nodes,omitempty"`
}

func (o InstallationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallationConfig struct{}"
	}

	return strings.Join([]string{"InstallationConfig", string(data)}, " ")
}
