package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransTemplateResponse Response Object
type UpdateTransTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTransTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateResponse", string(data)}, " ")
}
