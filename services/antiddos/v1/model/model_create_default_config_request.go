package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDefaultConfigRequest Request Object
type CreateDefaultConfigRequest struct {
	Body *DdosConfig `json:"body,omitempty"`
}

func (o CreateDefaultConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateDefaultConfigRequest", string(data)}, " ")
}
