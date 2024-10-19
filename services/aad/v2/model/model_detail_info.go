package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DetailInfo struct {
	Response *PageRespInfo `json:"response,omitempty"`
}

func (o DetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailInfo struct{}"
	}

	return strings.Join([]string{"DetailInfo", string(data)}, " ")
}
