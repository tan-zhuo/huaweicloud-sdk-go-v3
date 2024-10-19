package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDigitalAssetResponse Response Object
type CreateDigitalAssetResponse struct {

	// 数字资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDigitalAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDigitalAssetResponse struct{}"
	}

	return strings.Join([]string{"CreateDigitalAssetResponse", string(data)}, " ")
}
