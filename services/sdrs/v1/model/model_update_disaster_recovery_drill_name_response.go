package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterRecoveryDrillNameResponse Response Object
type UpdateDisasterRecoveryDrillNameResponse struct {
	DisasterRecoveryDrill *ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drill,omitempty"`
	HttpStatusCode        int                              `json:"-"`
}

func (o UpdateDisasterRecoveryDrillNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameResponse", string(data)}, " ")
}
