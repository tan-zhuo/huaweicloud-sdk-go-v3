package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectExtentionByIdCardImageResponse Response Object
type DetectExtentionByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result *IvsExtentionByIdCardImageResponseBodyResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectExtentionByIdCardImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageResponse", string(data)}, " ")
}
