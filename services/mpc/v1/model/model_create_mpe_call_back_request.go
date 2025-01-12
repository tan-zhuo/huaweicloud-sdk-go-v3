package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMpeCallBackRequest Request Object
type CreateMpeCallBackRequest struct {
	Body *MpeCallBackReq `json:"body,omitempty"`
}

func (o CreateMpeCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackRequest struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackRequest", string(data)}, " ")
}
