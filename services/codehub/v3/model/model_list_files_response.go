package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFilesResponse Response Object
type ListFilesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *FilesResponseInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesResponse struct{}"
	}

	return strings.Join([]string{"ListFilesResponse", string(data)}, " ")
}
