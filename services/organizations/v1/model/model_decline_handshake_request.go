package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeclineHandshakeRequest Request Object
type DeclineHandshakeRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 邀请的唯一标识符（ID）。账号在发起邀请时创建ID。
	HandshakeId string `json:"handshake_id"`
}

func (o DeclineHandshakeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeclineHandshakeRequest struct{}"
	}

	return strings.Join([]string{"DeclineHandshakeRequest", string(data)}, " ")
}
