package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnectionRequestBody 创建云连接实例的请求体。
type CreateCloudConnectionRequestBody struct {
	CloudConnection *CreateCloudConnection `json:"cloud_connection"`
}

func (o CreateCloudConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudConnectionRequestBody", string(data)}, " ")
}
