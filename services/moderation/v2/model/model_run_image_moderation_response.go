package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageModerationResponse Response Object
type RunImageModerationResponse struct {
	Result         *ImageDetectionResultBody `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RunImageModerationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageModerationResponse struct{}"
	}

	return strings.Join([]string{"RunImageModerationResponse", string(data)}, " ")
}
