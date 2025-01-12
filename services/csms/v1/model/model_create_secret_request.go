package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretRequest Request Object
type CreateSecretRequest struct {
	Body *CreateSecretRequestBody `json:"body,omitempty"`
}

func (o CreateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretRequest", string(data)}, " ")
}
