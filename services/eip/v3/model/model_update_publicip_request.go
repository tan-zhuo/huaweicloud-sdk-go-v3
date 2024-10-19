package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicipRequest Request Object
type UpdatePublicipRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id"`

	Body *UpdatePublicipsRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicipRequest", string(data)}, " ")
}
