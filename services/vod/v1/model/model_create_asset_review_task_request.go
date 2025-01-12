package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetReviewTaskRequest Request Object
type CreateAssetReviewTaskRequest struct {
	Body *AssetReviewReq `json:"body,omitempty"`
}

func (o CreateAssetReviewTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetReviewTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetReviewTaskRequest", string(data)}, " ")
}
