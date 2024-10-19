package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostResponse Response Object
type DeleteHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostResponse", string(data)}, " ")
}
