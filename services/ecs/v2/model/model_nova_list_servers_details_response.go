package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListServersDetailsResponse Response Object
type NovaListServersDetailsResponse struct {

	// 查询云服务器信息列表。
	Servers *[]NovaServer `json:"servers,omitempty"`

	// 分页查询时，查询下一页数据链接。
	ServersLinks   *[]PageLink `json:"servers_links,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o NovaListServersDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServersDetailsResponse struct{}"
	}

	return strings.Join([]string{"NovaListServersDetailsResponse", string(data)}, " ")
}
