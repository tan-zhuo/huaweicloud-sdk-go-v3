package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateGroupRequest Request Object
type UpdateTemplateGroupRequest struct {
	Body *ModifyTransTemplateGroup `json:"body,omitempty"`
}

func (o UpdateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupRequest", string(data)}, " ")
}
