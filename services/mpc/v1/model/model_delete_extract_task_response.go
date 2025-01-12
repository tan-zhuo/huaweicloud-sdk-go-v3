package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExtractTaskResponse Response Object
type DeleteExtractTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExtractTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteExtractTaskResponse", string(data)}, " ")
}
