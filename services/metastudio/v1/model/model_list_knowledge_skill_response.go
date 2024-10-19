package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKnowledgeSkillResponse Response Object
type ListKnowledgeSkillResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 技能信息
	Data *[]KnowledgeSkillInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListKnowledgeSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKnowledgeSkillResponse struct{}"
	}

	return strings.Join([]string{"ListKnowledgeSkillResponse", string(data)}, " ")
}
