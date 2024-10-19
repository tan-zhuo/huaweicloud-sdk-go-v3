package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefMessageChannelRequest Request Object
type CreateIefMessageChannelRequest struct {
	Body *CreateIefMessageChannelRequestBody `json:"body,omitempty"`
}

func (o CreateIefMessageChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefMessageChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateIefMessageChannelRequest", string(data)}, " ")
}
