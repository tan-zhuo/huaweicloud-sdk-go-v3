package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterEipRequest struct {
	Spec *MasterEipRequestSpec `json:"spec"`
}

func (o MasterEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequest struct{}"
	}

	return strings.Join([]string{"MasterEipRequest", string(data)}, " ")
}
