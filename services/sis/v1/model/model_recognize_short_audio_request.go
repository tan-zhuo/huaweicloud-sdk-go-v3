package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeShortAudioRequest Request Object
type RecognizeShortAudioRequest struct {
	Body *PostShortAudioReq `json:"body,omitempty"`
}

func (o RecognizeShortAudioRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeShortAudioRequest struct{}"
	}

	return strings.Join([]string{"RecognizeShortAudioRequest", string(data)}, " ")
}
