package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandSimilarityGraphTaskRequest Request Object
type CreateDrugLigandSimilarityGraphTaskRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateLigandSimilarityGraphTaskReq `json:"body,omitempty"`
}

func (o CreateDrugLigandSimilarityGraphTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandSimilarityGraphTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandSimilarityGraphTaskRequest", string(data)}, " ")
}
