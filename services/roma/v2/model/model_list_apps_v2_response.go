package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsV2Response Response Object
type ListAppsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// APP列表
	Apps           *[]AppInfoWithBindNum `json:"apps,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAppsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsV2Response struct{}"
	}

	return strings.Join([]string{"ListAppsV2Response", string(data)}, " ")
}
