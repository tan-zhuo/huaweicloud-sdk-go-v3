package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardWidgetsRequest Request Object
type ListDashboardWidgetsRequest struct {

	// 监控看板id，以db开头，包含22个字母和数字例：db16564943172807wjOmoLyn'
	DashboardId string `json:"dashboard_id"`
}

func (o ListDashboardWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardWidgetsRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardWidgetsRequest", string(data)}, " ")
}
