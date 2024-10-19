package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugJobResponse Response Object
type DeleteDrugJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugJobResponse", string(data)}, " ")
}
