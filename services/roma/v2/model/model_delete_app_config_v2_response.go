package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppConfigV2Response Response Object
type DeleteAppConfigV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppConfigV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppConfigV2Response struct{}"
	}

	return strings.Join([]string{"DeleteAppConfigV2Response", string(data)}, " ")
}
