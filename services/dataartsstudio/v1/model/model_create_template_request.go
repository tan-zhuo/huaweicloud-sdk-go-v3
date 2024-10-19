package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateRequest Request Object
type CreateTemplateRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *TemplateRo `json:"body,omitempty"`
}

func (o CreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequest", string(data)}, " ")
}
