package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyRequestBody
type UpdateAgencyRequestBody struct {
	Agency *UpdateAgencyOption `json:"agency"`
}

func (o UpdateAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAgencyRequestBody", string(data)}, " ")
}
