package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicipsRequestBody 更新弹性公网IP的请求体
type UpdatePublicipsRequestBody struct {
	Publicip *UpdatePublicipOption `json:"publicip"`
}

func (o UpdatePublicipsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePublicipsRequestBody", string(data)}, " ")
}
