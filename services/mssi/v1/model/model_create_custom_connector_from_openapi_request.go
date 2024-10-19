package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomConnectorFromOpenapiRequest Request Object
type CreateCustomConnectorFromOpenapiRequest struct {
	Body *CustomConnectorInfoV2 `json:"body,omitempty"`
}

func (o CreateCustomConnectorFromOpenapiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomConnectorFromOpenapiRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomConnectorFromOpenapiRequest", string(data)}, " ")
}
