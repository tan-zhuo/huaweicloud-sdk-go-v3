package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSentenceEmbeddingRequest Request Object
type RunSentenceEmbeddingRequest struct {
	Body *PostSentenceEmbeddingReq `json:"body,omitempty"`
}

func (o RunSentenceEmbeddingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSentenceEmbeddingRequest struct{}"
	}

	return strings.Join([]string{"RunSentenceEmbeddingRequest", string(data)}, " ")
}
