package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyResponse Response Object
type DeleteAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgencyResponse", string(data)}, " ")
}
