package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeploymentRequest Request Object
type UpdateDeploymentRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	Body *UpdateDeployment `json:"body,omitempty"`
}

func (o UpdateDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentRequest", string(data)}, " ")
}
