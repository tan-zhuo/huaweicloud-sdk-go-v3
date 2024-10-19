package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedVpcsResponse Response Object
type ListProtectedVpcsResponse struct {
	Data           *VpcProtectsVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListProtectedVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedVpcsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedVpcsResponse", string(data)}, " ")
}
