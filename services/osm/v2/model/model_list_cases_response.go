package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCasesResponse Response Object
type ListCasesResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 工单列表
	IncidentInfoList *[]IncidentInfoV2 `json:"incident_info_list,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ListCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCasesResponse struct{}"
	}

	return strings.Join([]string{"ListCasesResponse", string(data)}, " ")
}
