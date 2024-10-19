package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLastPropertyValueRequest Request Object
type ShowLastPropertyValueRequest struct {

	// 资产ID
	AssetId string `json:"asset_id"`

	Body *LastAssetPropertyValueRequest `json:"body,omitempty"`
}

func (o ShowLastPropertyValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLastPropertyValueRequest struct{}"
	}

	return strings.Join([]string{"ShowLastPropertyValueRequest", string(data)}, " ")
}
