package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceActionResponse Response Object
type UpdateInstanceActionResponse struct {

	// Job ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceActionResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceActionResponse", string(data)}, " ")
}
