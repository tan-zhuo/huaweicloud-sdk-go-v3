package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagsResponse Response Object
type DeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTagsResponse", string(data)}, " ")
}
