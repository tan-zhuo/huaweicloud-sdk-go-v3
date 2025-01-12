package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateRequestThrottlingPolicyV2Response Response Object
type DisassociateRequestThrottlingPolicyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateRequestThrottlingPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateRequestThrottlingPolicyV2Response", string(data)}, " ")
}
