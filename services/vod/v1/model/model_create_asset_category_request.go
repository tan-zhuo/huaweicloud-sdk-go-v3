package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetCategoryRequest Request Object
type CreateAssetCategoryRequest struct {
	Body *CreateCategoryReq `json:"body,omitempty"`
}

func (o CreateAssetCategoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetCategoryRequest", string(data)}, " ")
}
