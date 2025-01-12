package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePortRequest Request Object
type CreatePortRequest struct {
	Body *CreatePortRequestBody `json:"body,omitempty"`
}

func (o CreatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortRequest struct{}"
	}

	return strings.Join([]string{"CreatePortRequest", string(data)}, " ")
}
