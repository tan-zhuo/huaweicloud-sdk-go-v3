package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterRecoveryDrillNameRequestBody 更新容灾演练名称请求体
type UpdateDisasterRecoveryDrillNameRequestBody struct {
	DisasterRecoveryDrill *UpdateDisasterRecoveryDrillNameRequestParams `json:"disaster_recovery_drill"`
}

func (o UpdateDisasterRecoveryDrillNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameRequestBody", string(data)}, " ")
}
