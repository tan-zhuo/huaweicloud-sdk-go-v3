package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopRedistributionResponse Response Object
type StopRedistributionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopRedistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopRedistributionResponse struct{}"
	}

	return strings.Join([]string{"StopRedistributionResponse", string(data)}, " ")
}
