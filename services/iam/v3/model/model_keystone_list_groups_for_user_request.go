package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListGroupsForUserRequest Request Object
type KeystoneListGroupsForUserRequest struct {

	// 待查询的IAM用户ID，获取方式请参见：[获取项目名称、项目ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。
	UserId string `json:"user_id"`
}

func (o KeystoneListGroupsForUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListGroupsForUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListGroupsForUserRequest", string(data)}, " ")
}
