package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicipResponse Response Object
type UpdatePublicipResponse struct {
	Publicip       *PublicipUpdateResp `json:"publicip,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicipResponse", string(data)}, " ")
}
