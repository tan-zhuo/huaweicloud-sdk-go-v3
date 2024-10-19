package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoAssociateRouteEnabled struct {
}

func (o AutoAssociateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoAssociateRouteEnabled struct{}"
	}

	return strings.Join([]string{"AutoAssociateRouteEnabled", string(data)}, " ")
}
