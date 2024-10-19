package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployApplicationResponse Response Object
type DeployApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeployApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeployApplicationResponse", string(data)}, " ")
}
