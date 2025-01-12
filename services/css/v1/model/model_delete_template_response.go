package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateResponse Response Object
type DeleteTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateResponse", string(data)}, " ")
}
