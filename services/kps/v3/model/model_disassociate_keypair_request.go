package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateKeypairRequest Request Object
type DisassociateKeypairRequest struct {
	Body *DisassociateKeypairRequestBody `json:"body,omitempty"`
}

func (o DisassociateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateKeypairRequest struct{}"
	}

	return strings.Join([]string{"DisassociateKeypairRequest", string(data)}, " ")
}
