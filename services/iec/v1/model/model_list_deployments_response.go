package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeploymentsResponse Response Object
type ListDeploymentsResponse struct {

	// 部署计划列表的总和。
	Count *int32 `json:"count,omitempty"`

	// 部署计划列表。
	Deployments    *[]Deployment `json:"deployments,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListDeploymentsResponse", string(data)}, " ")
}
