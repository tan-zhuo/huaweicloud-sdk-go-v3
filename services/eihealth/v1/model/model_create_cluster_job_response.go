package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterJobResponse Response Object
type CreateClusterJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 限制的并发量
	LimitConcurrency *int32 `json:"limit_concurrency,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateClusterJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterJobResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterJobResponse", string(data)}, " ")
}
