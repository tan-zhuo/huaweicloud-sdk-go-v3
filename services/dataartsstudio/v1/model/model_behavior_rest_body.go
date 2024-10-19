package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BehaviorRestBody struct {
	Event string `json:"event"`

	Params []EventParam `json:"params"`
}

func (o BehaviorRestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorRestBody struct{}"
	}

	return strings.Join([]string{"BehaviorRestBody", string(data)}, " ")
}
