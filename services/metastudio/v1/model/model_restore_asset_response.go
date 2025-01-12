package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreAssetResponse Response Object
type RestoreAssetResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreAssetResponse struct{}"
	}

	return strings.Join([]string{"RestoreAssetResponse", string(data)}, " ")
}
