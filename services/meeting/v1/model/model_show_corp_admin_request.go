package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCorpAdminRequest Request Object
type ShowCorpAdminRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 企业用户帐号。 * 如果是帐号/密码鉴权方式，是指华为云会议帐号 * 如果是App ID鉴权方式，是指第三方User ID
	Account string `json:"account"`

	// 帐号类型。默认0。 * 0：华为云会议帐号。用于帐号/密码鉴权方式 * 1：第三方User ID，用于App ID鉴权方式
	AccountType *int32 `json:"accountType,omitempty"`
}

func (o ShowCorpAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpAdminRequest struct{}"
	}

	return strings.Join([]string{"ShowCorpAdminRequest", string(data)}, " ")
}
