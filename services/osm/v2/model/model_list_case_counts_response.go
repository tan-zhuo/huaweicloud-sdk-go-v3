package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaseCountsResponse Response Object
type ListCaseCountsResponse struct {

	// 状态数量统计列表
	IncidentStatusCounts *[]IncidentStatusCount `json:"incident_status_counts,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListCaseCountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseCountsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseCountsResponse", string(data)}, " ")
}
