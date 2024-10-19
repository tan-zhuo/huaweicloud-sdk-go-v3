package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetadataResponse Response Object
type CreateMetadataResponse struct {

	// 导入结果信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataResponse struct{}"
	}

	return strings.Join([]string{"CreateMetadataResponse", string(data)}, " ")
}
