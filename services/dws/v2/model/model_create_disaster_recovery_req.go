package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDisasterRecoveryReq struct {
	DisasterRecovery *CreateDisasterRecovery `json:"disaster_recovery"`
}

func (o CreateDisasterRecoveryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryReq struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryReq", string(data)}, " ")
}
