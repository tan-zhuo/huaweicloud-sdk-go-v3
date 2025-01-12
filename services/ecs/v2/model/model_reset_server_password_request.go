package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetServerPasswordRequest Request Object
type ResetServerPasswordRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ResetServerPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordRequest", string(data)}, " ")
}
