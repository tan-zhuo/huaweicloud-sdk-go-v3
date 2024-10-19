package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupsRequest Request Object
type ListAppGroupsRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`
}

func (o ListAppGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListAppGroupsRequest", string(data)}, " ")
}
