package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JsonNode struct {
}

func (o JsonNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JsonNode struct{}"
	}

	return strings.Join([]string{"JsonNode", string(data)}, " ")
}
