package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigprojectResponse Response Object
type DeleteMigprojectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigprojectResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectResponse", string(data)}, " ")
}
