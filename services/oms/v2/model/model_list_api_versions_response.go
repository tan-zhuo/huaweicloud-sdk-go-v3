package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiVersionsResponse Response Object
type ListApiVersionsResponse struct {

	// 版本信息列表。
	Versions       *[]Version `json:"versions,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
