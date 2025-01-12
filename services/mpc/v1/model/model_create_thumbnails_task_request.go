package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateThumbnailsTaskRequest Request Object
type CreateThumbnailsTaskRequest struct {
	Body *CreateThumbReq `json:"body,omitempty"`
}

func (o CreateThumbnailsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThumbnailsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateThumbnailsTaskRequest", string(data)}, " ")
}
