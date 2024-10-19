package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectExtentionByNameAndIdResponse Response Object
type DetectExtentionByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result *IvsExtentionByNameAndIdResponseBodyResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectExtentionByNameAndIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdResponse", string(data)}, " ")
}
