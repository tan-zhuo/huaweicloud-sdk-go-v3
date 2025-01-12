package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunReceptorPreprocessReq 受体预处理请求体
type RunReceptorPreprocessReq struct {
	File *ReceptorDrugFileReq `json:"file"`

	// 去除水分子
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除离子
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除配体分子
	RemoveLigand *bool `json:"remove_ligand,omitempty"`

	// 增加氢原子
	AddHydrogen *bool `json:"add_hydrogen,omitempty"`
}

func (o RunReceptorPreprocessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunReceptorPreprocessReq struct{}"
	}

	return strings.Join([]string{"RunReceptorPreprocessReq", string(data)}, " ")
}
