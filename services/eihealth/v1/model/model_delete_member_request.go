package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberRequest Request Object
type DeleteMemberRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 被移除项目的用户id
	UserId string `json:"user_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
