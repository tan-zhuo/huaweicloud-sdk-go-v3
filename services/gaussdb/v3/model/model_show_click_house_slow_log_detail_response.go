package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClickHouseSlowLogDetailResponse Response Object
type ShowClickHouseSlowLogDetailResponse struct {

	// 慢日志列表。
	SlowLogList    *[]ChSlowLogDetailResponseSlowLogList `json:"slow_log_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowClickHouseSlowLogDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseSlowLogDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClickHouseSlowLogDetailResponse", string(data)}, " ")
}
