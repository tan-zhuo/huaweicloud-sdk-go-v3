package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatePublicipsRequest Request Object
type AssociatePublicipsRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id"`

	Body *AssociatePublicipsRequestBody `json:"body,omitempty"`
}

func (o AssociatePublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePublicipsRequest struct{}"
	}

	return strings.Join([]string{"AssociatePublicipsRequest", string(data)}, " ")
}
