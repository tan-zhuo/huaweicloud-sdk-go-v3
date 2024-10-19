package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigRequest Request Object
type UpdateLtsConfigRequest struct {
	Body *UpdateLtsConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequest", string(data)}, " ")
}
