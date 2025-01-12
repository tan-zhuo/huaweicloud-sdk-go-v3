package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetMetaRequest Request Object
type UpdateAssetMetaRequest struct {
	Body *UpdateAssetMetaReq `json:"body,omitempty"`
}

func (o UpdateAssetMetaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetMetaRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetMetaRequest", string(data)}, " ")
}
