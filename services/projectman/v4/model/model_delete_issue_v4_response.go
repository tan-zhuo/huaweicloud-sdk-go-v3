package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIssueV4Response Response Object
type DeleteIssueV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIssueV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIssueV4Response struct{}"
	}

	return strings.Join([]string{"DeleteIssueV4Response", string(data)}, " ")
}
