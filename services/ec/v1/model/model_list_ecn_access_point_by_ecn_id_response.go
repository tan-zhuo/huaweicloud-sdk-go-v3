package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnAccessPointByEcnIdResponse Response Object
type ListEcnAccessPointByEcnIdResponse struct {
	AccessPoints   *[]AccessPoint `json:"access_points,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEcnAccessPointByEcnIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnAccessPointByEcnIdResponse struct{}"
	}

	return strings.Join([]string{"ListEcnAccessPointByEcnIdResponse", string(data)}, " ")
}
