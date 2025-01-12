package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisBindedToSignatureKeyV2Response Response Object
type ListApisBindedToSignatureKeyV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的签名密钥和API绑定关系列表
	Bindings       *[]SignApiBindingBase `json:"bindings,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListApisBindedToSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToSignatureKeyV2Response", string(data)}, " ")
}
