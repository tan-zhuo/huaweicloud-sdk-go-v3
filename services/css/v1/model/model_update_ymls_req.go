package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateYmlsReq struct {
	Edit *UpdateYmlsReqEdit `json:"edit"`
}

func (o UpdateYmlsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReq struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReq", string(data)}, " ")
}
