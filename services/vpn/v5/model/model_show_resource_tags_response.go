package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsResponse Response Object
type ShowResourceTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsResponse", string(data)}, " ")
}
