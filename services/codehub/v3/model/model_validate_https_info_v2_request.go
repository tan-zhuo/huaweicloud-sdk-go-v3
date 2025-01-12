package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateHttpsInfoV2Request Request Object
type ValidateHttpsInfoV2Request struct {

	// 用户iam_user_uuid
	IamUserUuid string `json:"iam_user_uuid"`

	Body *PasswordRequest `json:"body,omitempty"`
}

func (o ValidateHttpsInfoV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateHttpsInfoV2Request struct{}"
	}

	return strings.Join([]string{"ValidateHttpsInfoV2Request", string(data)}, " ")
}
