package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseUserInfoResponse Response Object
type UpdateDatabaseUserInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDatabaseUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseUserInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseUserInfoResponse", string(data)}, " ")
}
