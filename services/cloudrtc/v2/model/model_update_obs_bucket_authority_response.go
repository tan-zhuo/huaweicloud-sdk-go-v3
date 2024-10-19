package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsBucketAuthorityResponse Response Object
type UpdateObsBucketAuthorityResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateObsBucketAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsBucketAuthorityResponse struct{}"
	}

	return strings.Join([]string{"UpdateObsBucketAuthorityResponse", string(data)}, " ")
}
