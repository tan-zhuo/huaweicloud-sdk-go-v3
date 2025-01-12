package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterServerRequest Request Object
type RegisterServerRequest struct {
	Body *PostSourceServerBody `json:"body,omitempty"`
}

func (o RegisterServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerRequest struct{}"
	}

	return strings.Join([]string{"RegisterServerRequest", string(data)}, " ")
}
