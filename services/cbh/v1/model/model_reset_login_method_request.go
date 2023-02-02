package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetLoginMethodRequest struct {

	// Cbh Server Id
	ServerId string `json:"server_id"`
}

func (o ResetLoginMethodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLoginMethodRequest struct{}"
	}

	return strings.Join([]string{"ResetLoginMethodRequest", string(data)}, " ")
}