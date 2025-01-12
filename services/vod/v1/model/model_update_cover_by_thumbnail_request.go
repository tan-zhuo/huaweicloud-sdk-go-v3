package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCoverByThumbnailRequest Request Object
type UpdateCoverByThumbnailRequest struct {
	Body *UpdateCoverByThumbnailReq `json:"body,omitempty"`
}

func (o UpdateCoverByThumbnailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCoverByThumbnailRequest struct{}"
	}

	return strings.Join([]string{"UpdateCoverByThumbnailRequest", string(data)}, " ")
}
