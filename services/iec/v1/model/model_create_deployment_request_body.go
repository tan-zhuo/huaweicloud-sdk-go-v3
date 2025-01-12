package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentRequestBody 创建部署计划请求体
type CreateDeploymentRequestBody struct {
	Edgecloud *EdgeCloudOption `json:"edgecloud"`
}

func (o CreateDeploymentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDeploymentRequestBody", string(data)}, " ")
}
