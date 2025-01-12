package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicipsResponse Response Object
type ListPublicipsResponse struct {

	// 弹性公网IP对象
	Publicips      *[]PublicipShowResp `json:"publicips,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicipsResponse", string(data)}, " ")
}
