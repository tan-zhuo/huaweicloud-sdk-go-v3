package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRepoMembersResponse Response Object
type AddRepoMembersResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *[]CreateRepoMemberResult `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddRepoMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRepoMembersResponse struct{}"
	}

	return strings.Join([]string{"AddRepoMembersResponse", string(data)}, " ")
}
