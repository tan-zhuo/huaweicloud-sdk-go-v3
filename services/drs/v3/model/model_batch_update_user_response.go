package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateUserResponse Response Object
type BatchUpdateUserResponse struct {

	// 总数
	AllCounts *int32 `json:"all_counts,omitempty"`

	// 迁移用户信息
	Results        *[]QueryUserResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserResponse", string(data)}, " ")
}
