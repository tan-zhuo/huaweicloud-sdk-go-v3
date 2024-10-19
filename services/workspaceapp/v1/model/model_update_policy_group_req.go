package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyGroupReq struct {
	PolicyGroup *PolicyGroupForUpdate `json:"policy_group"`
}

func (o UpdatePolicyGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyGroupReq struct{}"
	}

	return strings.Join([]string{"UpdatePolicyGroupReq", string(data)}, " ")
}
