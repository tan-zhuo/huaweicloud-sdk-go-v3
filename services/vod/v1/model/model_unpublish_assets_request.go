package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishAssetsRequest Request Object
type UnpublishAssetsRequest struct {
	Body *PublishAssetReq `json:"body,omitempty"`
}

func (o UnpublishAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishAssetsRequest struct{}"
	}

	return strings.Join([]string{"UnpublishAssetsRequest", string(data)}, " ")
}
