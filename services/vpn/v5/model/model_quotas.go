package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quotas struct {
	Resources *[]QuotaInfo `json:"resources,omitempty"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
