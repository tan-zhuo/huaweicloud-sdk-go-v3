package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAssetFromObsResponse Response Object
type PublishAssetFromObsResponse struct {

	// 媒资ID
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishAssetFromObsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsResponse struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsResponse", string(data)}, " ")
}
