package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaProjectTagsRequest Request Object
type ShowKafkaProjectTagsRequest struct {
}

func (o ShowKafkaProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsRequest", string(data)}, " ")
}
