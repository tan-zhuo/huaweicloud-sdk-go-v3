package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStructTemplateResponse Response Object
type CreateStructTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateStructTemplateResponse", string(data)}, " ")
}
