package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLtsConfigsResponse Response Object
type CreateLtsConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"CreateLtsConfigsResponse", string(data)}, " ")
}
