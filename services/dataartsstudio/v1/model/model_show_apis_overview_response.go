package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApisOverviewResponse Response Object
type ShowApisOverviewResponse struct {

	// 已发布API量
	PublishNum *int32 `json:"publish_num,omitempty"`

	// 开发中API量
	DevelopingNum *int32 `json:"developing_num,omitempty"`

	// 申请量
	ApplyNum *int32 `json:"apply_num,omitempty"`

	// 调用总量
	CallNum *int32 `json:"call_num,omitempty"`

	// 成功调用量(取数成功)
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败调用量(取数失败)
	FailNum *int32 `json:"fail_num,omitempty"`

	// 合法调用量(通过校验)
	LegalNum *int32 `json:"legal_num,omitempty"`

	// 非法调用量(无法通过校验)
	IllegalNum     *int32 `json:"illegal_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowApisOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApisOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowApisOverviewResponse", string(data)}, " ")
}
