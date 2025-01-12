package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemporaryAccessKeyByAgencyRequest Request Object
type CreateTemporaryAccessKeyByAgencyRequest struct {
	Body *CreateTemporaryAccessKeyByAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateTemporaryAccessKeyByAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyRequest", string(data)}, " ")
}
