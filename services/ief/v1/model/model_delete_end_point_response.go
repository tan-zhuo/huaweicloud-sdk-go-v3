package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndPointResponse Response Object
type DeleteEndPointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndPointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndPointResponse", string(data)}, " ")
}
