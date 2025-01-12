package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateResourceResponse Response Object
type MigrateResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateResourceResponse struct{}"
	}

	return strings.Join([]string{"MigrateResourceResponse", string(data)}, " ")
}
