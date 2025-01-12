package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmRequest Request Object
type CreateAlarmRequest struct {
	Body *CreateAlarmRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequest", string(data)}, " ")
}
