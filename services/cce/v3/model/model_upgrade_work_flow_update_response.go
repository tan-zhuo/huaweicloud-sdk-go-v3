package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeWorkFlowUpdateResponse Response Object
type UpgradeWorkFlowUpdateResponse struct {

	// API类型，固定值“WorkFlowTask”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *WorkFlowSpec `json:"spec,omitempty"`

	Status         *WorkFlowStatus `json:"status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpgradeWorkFlowUpdateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeWorkFlowUpdateResponse struct{}"
	}

	return strings.Join([]string{"UpgradeWorkFlowUpdateResponse", string(data)}, " ")
}
