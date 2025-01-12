package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloseKibanaResponse Response Object
type UpdateCloseKibanaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCloseKibanaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloseKibanaResponse struct{}"
	}

	return strings.Join([]string{"UpdateCloseKibanaResponse", string(data)}, " ")
}
