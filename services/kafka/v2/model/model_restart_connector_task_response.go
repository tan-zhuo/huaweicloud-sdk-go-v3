package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartConnectorTaskResponse Response Object
type RestartConnectorTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestartConnectorTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartConnectorTaskResponse struct{}"
	}

	return strings.Join([]string{"RestartConnectorTaskResponse", string(data)}, " ")
}
