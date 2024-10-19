package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAccountRequest Request Object
type MoveAccountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	Body *MoveAccountReqBody `json:"body,omitempty"`
}

func (o MoveAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAccountRequest struct{}"
	}

	return strings.Join([]string{"MoveAccountRequest", string(data)}, " ")
}
