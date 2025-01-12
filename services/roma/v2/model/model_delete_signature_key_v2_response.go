package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSignatureKeyV2Response Response Object
type DeleteSignatureKeyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"DeleteSignatureKeyV2Response", string(data)}, " ")
}
