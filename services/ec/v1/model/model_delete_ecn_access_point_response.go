package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnAccessPointResponse Response Object
type DeleteEcnAccessPointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEcnAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnAccessPointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEcnAccessPointResponse", string(data)}, " ")
}
