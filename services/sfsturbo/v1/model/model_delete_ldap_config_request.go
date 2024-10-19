package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLdapConfigRequest Request Object
type DeleteLdapConfigRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`
}

func (o DeleteLdapConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLdapConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteLdapConfigRequest", string(data)}, " ")
}
