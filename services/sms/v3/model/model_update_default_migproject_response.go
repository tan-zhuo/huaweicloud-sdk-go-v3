package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefaultMigprojectResponse Response Object
type UpdateDefaultMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDefaultMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectResponse", string(data)}, " ")
}
