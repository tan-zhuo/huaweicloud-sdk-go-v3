package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSupplementdataRequest Request Object
type StopSupplementdataRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 补数据名称.
	InstanceName string `json:"instanceName"`
}

func (o StopSupplementdataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSupplementdataRequest struct{}"
	}

	return strings.Join([]string{"StopSupplementdataRequest", string(data)}, " ")
}
