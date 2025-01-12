package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStructTemplateResponse Response Object
type UpdateStructTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateStructTemplateResponse", string(data)}, " ")
}
