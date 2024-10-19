package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckImageSyncRequest Request Object
type BatchCheckImageSyncRequest struct {
	Body *ImageBatchSyncReq `json:"body,omitempty"`
}

func (o BatchCheckImageSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckImageSyncRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckImageSyncRequest", string(data)}, " ")
}
