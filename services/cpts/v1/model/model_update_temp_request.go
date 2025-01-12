package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTempRequest Request Object
type UpdateTempRequest struct {

	// 事务id
	TemplateId int32 `json:"template_id"`

	Body *UpdateTempRequestBody `json:"body,omitempty"`
}

func (o UpdateTempRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTempRequest struct{}"
	}

	return strings.Join([]string{"UpdateTempRequest", string(data)}, " ")
}
