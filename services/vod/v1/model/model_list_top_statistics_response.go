package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopStatisticsResponse Response Object
type ListTopStatisticsResponse struct {
	TopUrls        *[]TopUrl `json:"top_urls,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTopStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListTopStatisticsResponse", string(data)}, " ")
}
