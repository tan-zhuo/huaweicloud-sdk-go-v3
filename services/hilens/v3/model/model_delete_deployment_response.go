package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentResponse Response Object
type DeleteDeploymentResponse struct {

	// 部署Id
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentResponse", string(data)}, " ")
}
