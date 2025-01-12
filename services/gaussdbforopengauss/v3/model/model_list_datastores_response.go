package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastoresResponse Response Object
type ListDatastoresResponse struct {

	// 数据库版本号。
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatastoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoresResponse", string(data)}, " ")
}
