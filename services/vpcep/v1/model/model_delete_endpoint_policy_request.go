package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointPolicyRequest Request Object
type DeleteEndpointPolicyRequest struct {

	// 终端节点的ID。
	VpcEndpointId string `json:"vpc_endpoint_id"`

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`
}

func (o DeleteEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointPolicyRequest", string(data)}, " ")
}
