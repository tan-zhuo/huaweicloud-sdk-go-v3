package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterDetailsResponse Response Object
type ShowClusterDetailsResponse struct {
	Cluster        *Cluster `json:"cluster,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowClusterDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailsResponse", string(data)}, " ")
}
