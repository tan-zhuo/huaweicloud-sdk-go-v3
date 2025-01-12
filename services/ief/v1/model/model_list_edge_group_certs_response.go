package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeGroupCertsResponse Response Object
type ListEdgeGroupCertsResponse struct {
	Body           *[]EdgeGroupCertListResp `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListEdgeGroupCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeGroupCertsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeGroupCertsResponse", string(data)}, " ")
}
