package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFileResponse Response Object
type UpdateFileResponse struct {

	// 文件路径。
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFileResponse struct{}"
	}

	return strings.Join([]string{"UpdateFileResponse", string(data)}, " ")
}
