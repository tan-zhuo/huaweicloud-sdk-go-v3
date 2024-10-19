package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorisationRequest Request Object
type CreateAuthorisationRequest struct {
	Body *CreateAuthorisationRequestBody `json:"body,omitempty"`
}

func (o CreateAuthorisationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorisationRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorisationRequest", string(data)}, " ")
}
