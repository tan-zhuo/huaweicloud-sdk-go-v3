package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectStandardByNameAndIdResponse Response Object
type DetectStandardByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result *IvsStandardByNameAndIdResponseBodyResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectStandardByNameAndIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByNameAndIdResponse", string(data)}, " ")
}
