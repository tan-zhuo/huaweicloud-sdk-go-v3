package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationsResponse Response Object
type ListApplicationsResponse struct {

	// 应用总数。
	Count *int32 `json:"count,omitempty"`

	// 应用列表。
	Applications   *[]ApplicationView `json:"applications,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationsResponse", string(data)}, " ")
}
