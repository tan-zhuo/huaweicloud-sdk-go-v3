package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssuesRequest Request Object
type AssociateIssuesRequest struct {
	Body *AssociateIssuesRequestBody `json:"body,omitempty"`
}

func (o AssociateIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssuesRequest struct{}"
	}

	return strings.Join([]string{"AssociateIssuesRequest", string(data)}, " ")
}
