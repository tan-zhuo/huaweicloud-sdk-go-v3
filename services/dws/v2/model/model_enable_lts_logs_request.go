package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLtsLogsRequest Request Object
type EnableLtsLogsRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`
}

func (o EnableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
