package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationEndpointRequest Request Object
type CreateApplicationEndpointRequest struct {

	// Application的唯一资源标识，可通过[查询Application](smn_api_57004.xml)获取该标识。
	ApplicationUrn string `json:"application_urn"`

	Body *CreateApplicationEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateApplicationEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationEndpointRequest", string(data)}, " ")
}
