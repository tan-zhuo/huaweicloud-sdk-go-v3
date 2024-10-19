package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublishTemplateResponse Response Object
type UpdatePublishTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePublishTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublishTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublishTemplateResponse", string(data)}, " ")
}
