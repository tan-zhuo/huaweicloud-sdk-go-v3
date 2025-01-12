package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadIkThesaurusResponse Response Object
type CreateLoadIkThesaurusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateLoadIkThesaurusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadIkThesaurusResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadIkThesaurusResponse", string(data)}, " ")
}
