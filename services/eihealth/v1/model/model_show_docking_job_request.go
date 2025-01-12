package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDockingJobRequest Request Object
type ShowDockingJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowDockingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockingJobRequest struct{}"
	}

	return strings.Join([]string{"ShowDockingJobRequest", string(data)}, " ")
}
