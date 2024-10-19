package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppGroupAuthorizationRequest Request Object
type BatchDeleteAppGroupAuthorizationRequest struct {

	// 语言： - zh-cn：中文 - en-us：英文 - fr-fr: 法文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *AppGroupAuthorizeReq `json:"body,omitempty"`
}

func (o BatchDeleteAppGroupAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppGroupAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppGroupAuthorizationRequest", string(data)}, " ")
}
