package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageTaggingResponse Response Object
type RunImageTaggingResponse struct {
	Result         *ImageTaggingResponseResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RunImageTaggingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageTaggingResponse struct{}"
	}

	return strings.Join([]string{"RunImageTaggingResponse", string(data)}, " ")
}
