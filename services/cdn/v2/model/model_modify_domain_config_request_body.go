package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyDomainConfigRequestBody struct {
	Configs *Configs `json:"configs,omitempty"`
}

func (o ModifyDomainConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDomainConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDomainConfigRequestBody", string(data)}, " ")
}
