package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAssetTranscodeTaskRequest Request Object
type CancelAssetTranscodeTaskRequest struct {

	// 媒资ID。
	AssetId string `json:"asset_id"`
}

func (o CancelAssetTranscodeTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAssetTranscodeTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelAssetTranscodeTaskRequest", string(data)}, " ")
}
