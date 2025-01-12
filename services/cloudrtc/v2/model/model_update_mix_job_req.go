package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMixJobReq 合流任务参数
type UpdateMixJobReq struct {
	MixParam *UpdateMixParam `json:"mix_param"`
}

func (o UpdateMixJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMixJobReq struct{}"
	}

	return strings.Join([]string{"UpdateMixJobReq", string(data)}, " ")
}
