package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDatabasesResponse Response Object
type ListInstanceDatabasesResponse struct {

	// Redis实例数据库列表。
	Databases *[]string `json:"databases,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceDatabasesResponse", string(data)}, " ")
}
