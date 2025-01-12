package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateVaultPolicyResponse Response Object
type DisassociateVaultPolicyResponse struct {
	DissociatePolicy *VaultPolicyResp `json:"dissociate_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o DisassociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyResponse", string(data)}, " ")
}
