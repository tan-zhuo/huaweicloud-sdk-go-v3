package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyRequest Request Object
type CreateAgencyRequest struct {
	Body *AgencyRequestBody `json:"body,omitempty"`
}

func (o CreateAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequest", string(data)}, " ")
}
