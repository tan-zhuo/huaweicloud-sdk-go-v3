package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSimCardRequest struct {
	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取

	SimCardId int64 `json:"sim_card_id"`
}

func (o ShowSimCardRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSimCardRequest struct{}"
	}

	return strings.Join([]string{"ShowSimCardRequest", string(data)}, " ")
}