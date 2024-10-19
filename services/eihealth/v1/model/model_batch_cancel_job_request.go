package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCancelJobRequest Request Object
type BatchCancelJobRequest struct {

	// 是否强制停止作业
	XForce *bool `json:"X-Force,omitempty"`

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *BatchOperateJobReq `json:"body,omitempty"`
}

func (o BatchCancelJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCancelJobRequest struct{}"
	}

	return strings.Join([]string{"BatchCancelJobRequest", string(data)}, " ")
}
