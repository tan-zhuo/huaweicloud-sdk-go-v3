package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelCreateReqPrincipal struct {
	Iam *[]string `json:"IAM,omitempty"`
}

func (o ChannelCreateReqPrincipal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelCreateReqPrincipal struct{}"
	}

	return strings.Join([]string{"ChannelCreateReqPrincipal", string(data)}, " ")
}
