package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDocClassificationRequest Request Object
type RunDocClassificationRequest struct {
	Body *DocumentClassificationReq `json:"body,omitempty"`
}

func (o RunDocClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDocClassificationRequest struct{}"
	}

	return strings.Join([]string{"RunDocClassificationRequest", string(data)}, " ")
}
