package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDrugLigandToSmilesConversionRequest Request Object
type RunDrugLigandToSmilesConversionRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *DrugFile `json:"body,omitempty"`
}

func (o RunDrugLigandToSmilesConversionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDrugLigandToSmilesConversionRequest struct{}"
	}

	return strings.Join([]string{"RunDrugLigandToSmilesConversionRequest", string(data)}, " ")
}
