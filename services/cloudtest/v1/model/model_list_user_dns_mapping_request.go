package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserDnsMappingRequest Request Object
type ListUserDnsMappingRequest struct {

	// 工程ID
	ProjectId string `json:"project_id"`
}

func (o ListUserDnsMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserDnsMappingRequest struct{}"
	}

	return strings.Join([]string{"ListUserDnsMappingRequest", string(data)}, " ")
}
