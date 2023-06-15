package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCaTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCaTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCaTagResponse struct{}"
	}

	return strings.Join([]string{"CreateCaTagResponse", string(data)}, " ")
}
