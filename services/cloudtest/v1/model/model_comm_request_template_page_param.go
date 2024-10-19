package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestTemplatePageParam struct {
	Params *TemplatePageParam `json:"params,omitempty"`
}

func (o CommRequestTemplatePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestTemplatePageParam struct{}"
	}

	return strings.Join([]string{"CommRequestTemplatePageParam", string(data)}, " ")
}
