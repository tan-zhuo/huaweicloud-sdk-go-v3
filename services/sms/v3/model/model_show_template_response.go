package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateResponse Response Object
type ShowTemplateResponse struct {
	Template       *TemplateResponseBody `json:"template,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateResponse", string(data)}, " ")
}
