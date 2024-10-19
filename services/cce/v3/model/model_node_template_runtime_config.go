package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeTemplateRuntimeConfig struct {
	Runtime *NodeTemplateRuntimeConfigRuntime `json:"runtime,omitempty"`
}

func (o NodeTemplateRuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTemplateRuntimeConfig struct{}"
	}

	return strings.Join([]string{"NodeTemplateRuntimeConfig", string(data)}, " ")
}
