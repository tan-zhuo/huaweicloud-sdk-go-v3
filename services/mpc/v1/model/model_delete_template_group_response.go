package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateGroupResponse Response Object
type DeleteTemplateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupResponse", string(data)}, " ")
}
