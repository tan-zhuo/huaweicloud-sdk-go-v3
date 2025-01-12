package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTranscodingsTemplateResponse Response Object
type UpdateTranscodingsTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTranscodingsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTranscodingsTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTranscodingsTemplateResponse", string(data)}, " ")
}
