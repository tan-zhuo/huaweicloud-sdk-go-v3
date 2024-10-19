package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectSecretsTagsRequest Request Object
type ListProjectSecretsTagsRequest struct {
}

func (o ListProjectSecretsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectSecretsTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectSecretsTagsRequest", string(data)}, " ")
}
