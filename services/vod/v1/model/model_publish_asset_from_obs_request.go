package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAssetFromObsRequest Request Object
type PublishAssetFromObsRequest struct {
	Body *PublishAssetFromObsReq `json:"body,omitempty"`
}

func (o PublishAssetFromObsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsRequest struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsRequest", string(data)}, " ")
}
