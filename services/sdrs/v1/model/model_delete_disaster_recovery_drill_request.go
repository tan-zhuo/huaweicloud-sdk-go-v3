package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDisasterRecoveryDrillRequest Request Object
type DeleteDisasterRecoveryDrillRequest struct {

	// 容灾演练的ID。
	DisasterRecoveryDrillId string `json:"disaster_recovery_drill_id"`
}

func (o DeleteDisasterRecoveryDrillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryDrillRequest struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryDrillRequest", string(data)}, " ")
}
