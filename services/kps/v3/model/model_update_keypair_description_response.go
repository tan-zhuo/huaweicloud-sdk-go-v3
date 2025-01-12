package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeypairDescriptionResponse Response Object
type UpdateKeypairDescriptionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateKeypairDescriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairDescriptionResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeypairDescriptionResponse", string(data)}, " ")
}
