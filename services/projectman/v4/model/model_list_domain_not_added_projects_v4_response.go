package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainNotAddedProjectsV4Response Response Object
type ListDomainNotAddedProjectsV4Response struct {

	// 项目信息列表
	Projects *[]ListDomainNotAddedProjectsV4ResponseBodyProjects `json:"projects,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDomainNotAddedProjectsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainNotAddedProjectsV4Response struct{}"
	}

	return strings.Join([]string{"ListDomainNotAddedProjectsV4Response", string(data)}, " ")
}
