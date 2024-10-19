package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedisModifyDbUserPrivilegeRequest struct {
	Users *[]RedisModifyDbUserPrivilegeRequestBody `json:"users,omitempty"`
}

func (o RedisModifyDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisModifyDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"RedisModifyDbUserPrivilegeRequest", string(data)}, " ")
}
