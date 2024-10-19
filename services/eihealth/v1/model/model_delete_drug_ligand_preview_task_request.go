package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugLigandPreviewTaskRequest Request Object
type DeleteDrugLigandPreviewTaskRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 任务ID，通过创建任务接口取得。
	TaskId string `json:"task_id"`
}

func (o DeleteDrugLigandPreviewTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugLigandPreviewTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteDrugLigandPreviewTaskRequest", string(data)}, " ")
}
