package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstanceRequestBody struct {
	Source *RecoveryBackupSource `json:"source"`

	Target *RecoveryBackupTarget `json:"target"`
}

func (o RestoreInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequestBody", string(data)}, " ")
}
