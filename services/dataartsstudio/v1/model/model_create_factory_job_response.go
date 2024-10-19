package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryJobResponse Response Object
type CreateFactoryJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateFactoryJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryJobResponse struct{}"
	}

	return strings.Join([]string{"CreateFactoryJobResponse", string(data)}, " ")
}
