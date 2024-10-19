package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlLimitRulesResponse Response Object
type DeleteSqlLimitRulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSqlLimitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitRulesResponse", string(data)}, " ")
}
