package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVocabularyRequest Request Object
type DeleteVocabularyRequest struct {

	// 热词表id。
	VocabularyId string `json:"vocabulary_id"`
}

func (o DeleteVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVocabularyRequest struct{}"
	}

	return strings.Join([]string{"DeleteVocabularyRequest", string(data)}, " ")
}
