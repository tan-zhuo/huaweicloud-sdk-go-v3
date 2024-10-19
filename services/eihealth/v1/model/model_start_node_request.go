package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartNodeRequest Request Object
type StartNodeRequest struct {

	// 计算资源id
	Id string `json:"id"`
}

func (o StartNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNodeRequest struct{}"
	}

	return strings.Join([]string{"StartNodeRequest", string(data)}, " ")
}
