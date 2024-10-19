package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPackagesRequest Request Object
type ListPackagesRequest struct {
	Body *PackagesListRequestBody `json:"body,omitempty"`
}

func (o ListPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackagesRequest struct{}"
	}

	return strings.Join([]string{"ListPackagesRequest", string(data)}, " ")
}
