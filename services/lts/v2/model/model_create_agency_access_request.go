package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyAccessRequest Request Object
type CreateAgencyAccessRequest struct {
	Body *PreviewAgencyLogAccessReqListBody `json:"body,omitempty"`
}

func (o CreateAgencyAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyAccessRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyAccessRequest", string(data)}, " ")
}
