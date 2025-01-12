package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermanentAccessKeyRequest Request Object
type DeletePermanentAccessKeyRequest struct {

	// 待删除的指定AK。
	AccessKey string `json:"access_key"`
}

func (o DeletePermanentAccessKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"DeletePermanentAccessKeyRequest", string(data)}, " ")
}
