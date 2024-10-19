package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServersRequest Request Object
type DeleteServersRequest struct {
	Body *DeleteServersRequestBody `json:"body,omitempty"`
}

func (o DeleteServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersRequest struct{}"
	}

	return strings.Join([]string{"DeleteServersRequest", string(data)}, " ")
}
