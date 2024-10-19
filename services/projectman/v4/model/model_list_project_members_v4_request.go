package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMembersV4Request Request Object
type ListProjectMembersV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 分页索引，偏移量,offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量,每页最多显示100条
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectMembersV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMembersV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectMembersV4Request", string(data)}, " ")
}
