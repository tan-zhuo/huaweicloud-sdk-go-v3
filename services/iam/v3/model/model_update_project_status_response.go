package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectStatusResponse Response Object
type UpdateProjectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProjectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectStatusResponse", string(data)}, " ")
}
