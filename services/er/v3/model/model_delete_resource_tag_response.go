package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceTagResponse Response Object
type DeleteResourceTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagResponse", string(data)}, " ")
}
