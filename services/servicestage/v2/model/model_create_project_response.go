package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectResponse Response Object
type CreateProjectResponse struct {

	// 项目ID。
	Id *string `json:"id,omitempty"`

	// 项目名称。
	Name *string `json:"name,omitempty"`

	// 项目的clone url路径。
	CloneUrl       *string `json:"clone_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectResponse", string(data)}, " ")
}
