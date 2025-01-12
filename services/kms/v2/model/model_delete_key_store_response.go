package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeyStoreResponse Response Object
type DeleteKeyStoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeyStoreResponse", string(data)}, " ")
}
