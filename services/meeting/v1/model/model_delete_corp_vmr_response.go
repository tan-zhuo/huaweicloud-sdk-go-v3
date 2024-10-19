package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCorpVmrResponse Response Object
type DeleteCorpVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCorpVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCorpVmrResponse struct{}"
	}

	return strings.Join([]string{"DeleteCorpVmrResponse", string(data)}, " ")
}
