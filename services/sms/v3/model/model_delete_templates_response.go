package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplatesResponse Response Object
type DeleteTemplatesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesResponse", string(data)}, " ")
}
