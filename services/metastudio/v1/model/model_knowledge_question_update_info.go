package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeQuestionUpdateInfo 修改知识库问法请求。
type KnowledgeQuestionUpdateInfo struct {

	// 问法ID。
	QuestionId string `json:"question_id"`

	// 问法。
	Question string `json:"question"`
}

func (o KnowledgeQuestionUpdateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeQuestionUpdateInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeQuestionUpdateInfo", string(data)}, " ")
}
