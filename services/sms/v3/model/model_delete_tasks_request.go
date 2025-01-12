package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTasksRequest Request Object
type DeleteTasksRequest struct {
	Body *DeleteTasksReq `json:"body,omitempty"`
}

func (o DeleteTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteTasksRequest", string(data)}, " ")
}
