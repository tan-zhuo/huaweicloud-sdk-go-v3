package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeReceptorPocketReq 受体口袋识别请求体
type RecognizeReceptorPocketReq struct {
	Mode *RecognizeReceptorPocketMode `json:"mode"`

	ReceptorFile *ReceptorDrugFileReq `json:"receptor_file"`

	LigandFile *DrugFile `json:"ligand_file,omitempty"`

	// 残基列表，当识别模式为残基时提供
	Residues *[]string `json:"residues,omitempty"`
}

func (o RecognizeReceptorPocketReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeReceptorPocketReq struct{}"
	}

	return strings.Join([]string{"RecognizeReceptorPocketReq", string(data)}, " ")
}
