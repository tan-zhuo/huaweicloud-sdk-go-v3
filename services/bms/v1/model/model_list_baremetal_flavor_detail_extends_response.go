package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaremetalFlavorDetailExtendsResponse Response Object
type ListBaremetalFlavorDetailExtendsResponse struct {

	// 裸金属服务器规格列表，详情请参见表2 flavors数据结构说明。
	Flavors        *[]FlavorsResp `json:"flavors,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListBaremetalFlavorDetailExtendsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaremetalFlavorDetailExtendsResponse struct{}"
	}

	return strings.Join([]string{"ListBaremetalFlavorDetailExtendsResponse", string(data)}, " ")
}
