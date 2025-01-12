package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmFileUploadResponse Response Object
type ConfirmFileUploadResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmFileUploadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmFileUploadResponse struct{}"
	}

	return strings.Join([]string{"ConfirmFileUploadResponse", string(data)}, " ")
}
