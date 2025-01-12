package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppVersionsResponse Response Object
type ListAppVersionsResponse struct {

	// app详情
	Versions *[]AppVersionDetail `json:"versions,omitempty"`

	// 满足条件的应用版本个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListAppVersionsResponse", string(data)}, " ")
}
