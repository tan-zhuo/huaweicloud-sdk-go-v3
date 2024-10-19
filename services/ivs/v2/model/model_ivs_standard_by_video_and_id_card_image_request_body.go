package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsStandardByVideoAndIdCardImageRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsStandardByVideoAndIdCardImageRequestBodyData `json:"data"`
}

func (o IvsStandardByVideoAndIdCardImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByVideoAndIdCardImageRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByVideoAndIdCardImageRequestBody", string(data)}, " ")
}
