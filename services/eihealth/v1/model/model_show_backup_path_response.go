package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPathResponse Response Object
type ShowBackupPathResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowBackupPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPathResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPathResponse", string(data)}, " ")
}
