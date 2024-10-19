package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDatabaseOmUserActionRequest Request Object
type ExecuteDatabaseOmUserActionRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *DatabaseOmUserActionReq `json:"body,omitempty"`
}

func (o ExecuteDatabaseOmUserActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDatabaseOmUserActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDatabaseOmUserActionRequest", string(data)}, " ")
}
