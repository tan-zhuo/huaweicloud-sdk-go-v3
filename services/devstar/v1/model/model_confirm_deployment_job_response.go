package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmDeploymentJobResponse Response Object
type ConfirmDeploymentJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ConfirmDeploymentJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDeploymentJobResponse struct{}"
	}

	return strings.Join([]string{"ConfirmDeploymentJobResponse", string(data)}, " ")
}
