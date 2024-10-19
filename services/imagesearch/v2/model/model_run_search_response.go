package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSearchResponse Response Object
type RunSearchResponse struct {

	// 搜索完成返回success。
	Result *string `json:"result,omitempty"`

	Data           *SearchRestInfo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RunSearchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchResponse struct{}"
	}

	return strings.Join([]string{"RunSearchResponse", string(data)}, " ")
}
