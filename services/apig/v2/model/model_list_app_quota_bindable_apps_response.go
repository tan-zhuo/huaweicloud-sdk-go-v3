package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppQuotaBindableAppsResponse Response Object
type ListAppQuotaBindableAppsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 凭据列表
	Apps           *[]AppQuotaAppInfo `json:"apps,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppQuotaBindableAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuotaBindableAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppQuotaBindableAppsResponse", string(data)}, " ")
}
