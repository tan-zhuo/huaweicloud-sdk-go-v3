package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageTagResponse Response Object
type AddImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagResponse struct{}"
	}

	return strings.Join([]string{"AddImageTagResponse", string(data)}, " ")
}
