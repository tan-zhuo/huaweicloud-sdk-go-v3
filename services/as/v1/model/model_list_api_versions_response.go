package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiVersionsResponse Response Object
type ListApiVersionsResponse struct {

	// 弹性伸缩API版本信息。
	Versions       *[]VersionInfo `json:"versions,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
