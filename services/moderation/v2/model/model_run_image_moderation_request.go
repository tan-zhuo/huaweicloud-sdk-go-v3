package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageModerationRequest Request Object
type RunImageModerationRequest struct {
	Body *ImageDetectionReq `json:"body,omitempty"`
}

func (o RunImageModerationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageModerationRequest struct{}"
	}

	return strings.Join([]string{"RunImageModerationRequest", string(data)}, " ")
}
