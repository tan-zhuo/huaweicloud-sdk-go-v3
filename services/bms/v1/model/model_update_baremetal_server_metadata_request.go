package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerMetadataRequest Request Object
type UpdateBaremetalServerMetadataRequest struct {

	// 裸金属服务器ID
	ServerId string `json:"server_id"`

	Body *UpdateBaremetalServerMetadataReq `json:"body,omitempty"`
}

func (o UpdateBaremetalServerMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataRequest struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataRequest", string(data)}, " ")
}
