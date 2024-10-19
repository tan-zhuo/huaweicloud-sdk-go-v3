package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAtomicIndexesResponse Response Object
type SearchAtomicIndexesResponse struct {
	Data           *AtomicIndexVoSearchResultData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o SearchAtomicIndexesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAtomicIndexesResponse struct{}"
	}

	return strings.Join([]string{"SearchAtomicIndexesResponse", string(data)}, " ")
}
