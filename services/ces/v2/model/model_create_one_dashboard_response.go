package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOneDashboardResponse Response Object
type CreateOneDashboardResponse struct {

	// 监控看板id
	DashboardId    *string `json:"dashboard_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOneDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOneDashboardResponse struct{}"
	}

	return strings.Join([]string{"CreateOneDashboardResponse", string(data)}, " ")
}
