package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMultiPictureResponse Response Object
type SetMultiPictureResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetMultiPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMultiPictureResponse struct{}"
	}

	return strings.Join([]string{"SetMultiPictureResponse", string(data)}, " ")
}
