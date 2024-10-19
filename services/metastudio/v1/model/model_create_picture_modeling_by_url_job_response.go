package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePictureModelingByUrlJobResponse Response Object
type CreatePictureModelingByUrlJobResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePictureModelingByUrlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePictureModelingByUrlJobResponse struct{}"
	}

	return strings.Join([]string{"CreatePictureModelingByUrlJobResponse", string(data)}, " ")
}
