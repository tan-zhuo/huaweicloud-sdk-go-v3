package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachApiFromPluginResponse Response Object
type DetachApiFromPluginResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachApiFromPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachApiFromPluginResponse struct{}"
	}

	return strings.Join([]string{"DetachApiFromPluginResponse", string(data)}, " ")
}
