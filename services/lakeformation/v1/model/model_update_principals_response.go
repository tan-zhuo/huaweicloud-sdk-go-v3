package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrincipalsResponse Response Object
type UpdatePrincipalsResponse struct {
	Body           *[]Principal `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdatePrincipalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrincipalsResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrincipalsResponse", string(data)}, " ")
}
