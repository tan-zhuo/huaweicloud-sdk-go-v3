package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIkThesaurusResponse Response Object
type DeleteIkThesaurusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIkThesaurusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIkThesaurusResponse struct{}"
	}

	return strings.Join([]string{"DeleteIkThesaurusResponse", string(data)}, " ")
}
