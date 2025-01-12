package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkSpaceRequest Request Object
type ShowWorkSpaceRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowWorkSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkSpaceRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkSpaceRequest", string(data)}, " ")
}
