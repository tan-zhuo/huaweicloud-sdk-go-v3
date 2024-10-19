package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteComposeVideoResponse Response Object
type ExecuteComposeVideoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteComposeVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteComposeVideoResponse struct{}"
	}

	return strings.Join([]string{"ExecuteComposeVideoResponse", string(data)}, " ")
}
