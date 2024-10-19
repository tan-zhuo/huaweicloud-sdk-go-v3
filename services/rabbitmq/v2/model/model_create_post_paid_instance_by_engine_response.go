package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidInstanceByEngineResponse Response Object
type CreatePostPaidInstanceByEngineResponse struct {

	// 实例ID。
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidInstanceByEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceByEngineResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceByEngineResponse", string(data)}, " ")
}
