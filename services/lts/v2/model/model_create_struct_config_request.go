package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStructConfigRequest Request Object
type CreateStructConfigRequest struct {
	Body *StructConfig `json:"body,omitempty"`
}

func (o CreateStructConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateStructConfigRequest", string(data)}, " ")
}
