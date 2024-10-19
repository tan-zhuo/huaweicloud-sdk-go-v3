package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTrashesRequest Request Object
type BatchDeleteTrashesRequest struct {
	Body *[]TrashArtifactModelForDelete `json:"body,omitempty"`
}

func (o BatchDeleteTrashesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTrashesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTrashesRequest", string(data)}, " ")
}
