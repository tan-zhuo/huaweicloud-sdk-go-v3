package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveVaultResourceResponse Response Object
type RemoveVaultResourceResponse struct {

	// 移除的资源ID
	RemoveResourceIds *[]string `json:"remove_resource_ids,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o RemoveVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"RemoveVaultResourceResponse", string(data)}, " ")
}
