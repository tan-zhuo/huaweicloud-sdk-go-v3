package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePredefineTagsResponse Response Object
type UpdatePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePredefineTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"UpdatePredefineTagsResponse", string(data)}, " ")
}
