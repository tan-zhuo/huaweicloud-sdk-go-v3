package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEncryptTaskRequest Request Object
type CreateEncryptTaskRequest struct {
	Body *CreateEncryptReq `json:"body,omitempty"`
}

func (o CreateEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskRequest", string(data)}, " ")
}
