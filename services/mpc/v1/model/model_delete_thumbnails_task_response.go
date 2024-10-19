package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteThumbnailsTaskResponse Response Object
type DeleteThumbnailsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteThumbnailsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsTaskResponse", string(data)}, " ")
}
