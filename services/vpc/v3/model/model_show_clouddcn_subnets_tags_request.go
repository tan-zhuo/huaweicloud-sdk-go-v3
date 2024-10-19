package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClouddcnSubnetsTagsRequest Request Object
type ShowClouddcnSubnetsTagsRequest struct {

	// Clouddcn子网的id
	ResourceId string `json:"resource_id"`
}

func (o ShowClouddcnSubnetsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClouddcnSubnetsTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowClouddcnSubnetsTagsRequest", string(data)}, " ")
}
