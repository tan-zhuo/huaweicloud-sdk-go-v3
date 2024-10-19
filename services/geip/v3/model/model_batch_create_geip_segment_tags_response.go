package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGeipSegmentTagsResponse Response Object
type BatchCreateGeipSegmentTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateGeipSegmentTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGeipSegmentTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateGeipSegmentTagsResponse", string(data)}, " ")
}
