package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSummaryDomainRequest Request Object
type RunSummaryDomainRequest struct {
	Body *SummaryDomainReq `json:"body,omitempty"`
}

func (o RunSummaryDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSummaryDomainRequest struct{}"
	}

	return strings.Join([]string{"RunSummaryDomainRequest", string(data)}, " ")
}
