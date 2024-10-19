package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthCheckRequestBody create Health Check request
type CreateHealthCheckRequestBody struct {
	HealthCheck *CreateHealthCheckOption `json:"health_check"`
}

func (o CreateHealthCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthCheckRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthCheckRequestBody", string(data)}, " ")
}
