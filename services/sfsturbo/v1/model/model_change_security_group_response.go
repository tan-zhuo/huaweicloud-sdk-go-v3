package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityGroupResponse Response Object
type ChangeSecurityGroupResponse struct {

	// SFS Turbo文件系统ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupResponse", string(data)}, " ")
}
