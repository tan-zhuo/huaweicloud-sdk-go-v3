package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PasswordPolicyOption
type PasswordPolicyOption struct {

	// 同一字符连续出现的最大次数，取值范围[0,32]。
	MaximumConsecutiveIdenticalChars *int32 `json:"maximum_consecutive_identical_chars,omitempty"`

	// 密码最短使用时间(分钟)，取值范围[0,1440]。
	MinimumPasswordAge *int32 `json:"minimum_password_age,omitempty"`

	// 密码最小字符数，取值范围[6,32]。
	MinimumPasswordLength *int32 `json:"minimum_password_length,omitempty"`

	// 密码不能与历史密码重复次数，取值范围[0,10]。
	NumberOfRecentPasswordsDisallowed *int32 `json:"number_of_recent_passwords_disallowed,omitempty"`

	// 密码是否可以是用户名或用户名的反序。
	PasswordNotUsernameOrInvert *bool `json:"password_not_username_or_invert,omitempty"`

	// 密码有效期（天），取值范围[0,180]，设置0表示关闭该策略。
	PasswordValidityPeriod *int32 `json:"password_validity_period,omitempty"`

	// 至少包含字符种类的个数，取值区间[2,4]。
	PasswordCharCombination *int32 `json:"password_char_combination,omitempty"`
}

func (o PasswordPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordPolicyOption struct{}"
	}

	return strings.Join([]string{"PasswordPolicyOption", string(data)}, " ")
}
