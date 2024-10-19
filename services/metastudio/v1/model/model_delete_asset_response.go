package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetResponse Response Object
type DeleteAssetResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetResponse", string(data)}, " ")
}
