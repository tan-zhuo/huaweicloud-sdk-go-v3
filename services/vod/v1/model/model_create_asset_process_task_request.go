package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetProcessTaskRequest Request Object
type CreateAssetProcessTaskRequest struct {
	Body *AssetProcessReq `json:"body,omitempty"`
}

func (o CreateAssetProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskRequest", string(data)}, " ")
}
