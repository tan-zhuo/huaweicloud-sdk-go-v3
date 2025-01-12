package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStructTemplateResponse Response Object
type DeleteStructTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteStructTemplateResponse", string(data)}, " ")
}
