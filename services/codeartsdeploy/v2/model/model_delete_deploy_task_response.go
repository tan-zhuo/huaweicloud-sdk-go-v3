package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeployTaskResponse Response Object
type DeleteDeployTaskResponse struct {

	// 部署任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeployTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeployTaskResponse", string(data)}, " ")
}
