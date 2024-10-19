package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtomicIndexVoSearchResultData 返回数据。
type AtomicIndexVoSearchResultData struct {
	Value *AtomicIndexVoSearchResultDataValue `json:"value,omitempty"`
}

func (o AtomicIndexVoSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicIndexVoSearchResultData struct{}"
	}

	return strings.Join([]string{"AtomicIndexVoSearchResultData", string(data)}, " ")
}
