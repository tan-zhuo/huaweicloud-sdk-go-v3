package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVaultTagsResponse Response Object
type CreateVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVaultTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsResponse", string(data)}, " ")
}
