package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterTagRequest Request Object
type CreateClusterTagRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreateClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterTagRequest", string(data)}, " ")
}
