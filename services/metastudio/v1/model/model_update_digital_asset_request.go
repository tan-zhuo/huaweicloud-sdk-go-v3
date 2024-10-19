package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDigitalAssetRequest Request Object
type UpdateDigitalAssetRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 资产ID。
	AssetId string `json:"asset_id"`

	Body *UpdateDigitalAssetRequestBody `json:"body,omitempty"`
}

func (o UpdateDigitalAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDigitalAssetRequest struct{}"
	}

	return strings.Join([]string{"UpdateDigitalAssetRequest", string(data)}, " ")
}
