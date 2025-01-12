package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePoolResponse Response Object
type DeletePoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolResponse struct{}"
	}

	return strings.Join([]string{"DeletePoolResponse", string(data)}, " ")
}
