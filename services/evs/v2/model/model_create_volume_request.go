package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVolumeRequest Request Object
type CreateVolumeRequest struct {
	Body *CreateVolumeRequestBody `json:"body,omitempty"`
}

func (o CreateVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeRequest struct{}"
	}

	return strings.Join([]string{"CreateVolumeRequest", string(data)}, " ")
}
