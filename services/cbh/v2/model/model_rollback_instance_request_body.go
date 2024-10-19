package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollbackInstanceRequestBody struct {

	// 需要回退的实例id
	ServerId string `json:"server_id"`
}

func (o RollbackInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RollbackInstanceRequestBody", string(data)}, " ")
}
