package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouchDbInfo struct {

	// couchDB用户名称
	User *string `json:"user,omitempty"`
}

func (o CouchDbInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouchDbInfo struct{}"
	}

	return strings.Join([]string{"CouchDbInfo", string(data)}, " ")
}
