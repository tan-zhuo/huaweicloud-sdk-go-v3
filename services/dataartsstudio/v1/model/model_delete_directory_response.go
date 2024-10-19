package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDirectoryResponse Response Object
type DeleteDirectoryResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteDirectoryResponse", string(data)}, " ")
}
