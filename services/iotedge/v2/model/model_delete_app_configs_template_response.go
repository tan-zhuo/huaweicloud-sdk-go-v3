package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppConfigsTemplateResponse Response Object
type DeleteAppConfigsTemplateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAppConfigsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppConfigsTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppConfigsTemplateResponse", string(data)}, " ")
}
