package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetModelResponse Response Object
type DeleteAssetModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetModelResponse", string(data)}, " ")
}
