package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
