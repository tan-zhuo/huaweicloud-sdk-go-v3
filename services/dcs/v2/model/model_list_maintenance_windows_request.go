package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMaintenanceWindowsRequest Request Object
type ListMaintenanceWindowsRequest struct {
}

func (o ListMaintenanceWindowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMaintenanceWindowsRequest struct{}"
	}

	return strings.Join([]string{"ListMaintenanceWindowsRequest", string(data)}, " ")
}
