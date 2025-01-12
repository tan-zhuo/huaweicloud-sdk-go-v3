package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicationResponse Response Object
type UpdatePublicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicationResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicationResponse", string(data)}, " ")
}
