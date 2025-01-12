package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDbUserResponse Response Object
type RegisterDbUserResponse struct {

	// 数据库用户ID
	DbUserId       *string `json:"db_user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDbUserResponse struct{}"
	}

	return strings.Join([]string{"RegisterDbUserResponse", string(data)}, " ")
}
