package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeVinResponse Response Object
type RecognizeVinResponse struct {
	Result *VinResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeVinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeVinResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVinResponse", string(data)}, " ")
}
