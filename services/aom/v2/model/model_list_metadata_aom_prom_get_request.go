package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetadataAomPromGetRequest Request Object
type ListMetadataAomPromGetRequest struct {
}

func (o ListMetadataAomPromGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadataAomPromGetRequest struct{}"
	}

	return strings.Join([]string{"ListMetadataAomPromGetRequest", string(data)}, " ")
}
