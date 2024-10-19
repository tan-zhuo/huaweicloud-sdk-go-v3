package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResultRequest Request Object
type ShowInstanceResultRequest struct {

	// projectId
	InstanceId string `json:"instance_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowInstanceResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResultRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceResultRequest", string(data)}, " ")
}
