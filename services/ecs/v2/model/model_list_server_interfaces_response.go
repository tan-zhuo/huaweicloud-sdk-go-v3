package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerInterfacesResponse Response Object
type ListServerInterfacesResponse struct {
	AttachableQuantity *InterfaceAttachableQuantity `json:"attachableQuantity,omitempty"`

	// 云服务器网卡信息列表
	InterfaceAttachments *[]InterfaceAttachment `json:"interfaceAttachments,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListServerInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListServerInterfacesResponse", string(data)}, " ")
}
