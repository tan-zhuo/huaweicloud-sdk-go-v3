package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlDatabaseUserRequest Request Object
type DeleteGaussMySqlDatabaseUserRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseUserRequest `json:"body,omitempty"`
}

func (o DeleteGaussMySqlDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlDatabaseUserRequest", string(data)}, " ")
}
