package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmImageUploadRequest Request Object
type ConfirmImageUploadRequest struct {
	Body *ConfirmImageUploadReq `json:"body,omitempty"`
}

func (o ConfirmImageUploadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmImageUploadRequest struct{}"
	}

	return strings.Join([]string{"ConfirmImageUploadRequest", string(data)}, " ")
}
