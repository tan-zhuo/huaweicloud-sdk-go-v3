package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstancesSessionRespondBodySessions struct {

	// 客户端的地址和端口。
	Addr *string `json:"addr,omitempty"`

	// 会话id。
	Id *string `json:"id,omitempty"`

	// 连接名。
	Name *string `json:"name,omitempty"`

	// 最近一次执行的命令。
	Cmd *string `json:"cmd,omitempty"`

	// 以秒计算的已连接时长。
	Age *string `json:"age,omitempty"`

	// 以秒计算的空闲时长。
	Idle *string `json:"idle,omitempty"`

	// 该客户端正在使用的数据库 ID。
	Db *string `json:"db,omitempty"`

	// 套接字所使用的文件描述符。
	Fd *string `json:"fd,omitempty"`

	// 已订阅频道的数量。
	Sub *string `json:"sub,omitempty"`

	// 已订阅模式的数量。
	Psub *string `json:"psub,omitempty"`

	// 在事务中被执行的命令数量。
	Multi *string `json:"multi,omitempty"`
}

func (o ListInstancesSessionRespondBodySessions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionRespondBodySessions struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionRespondBodySessions", string(data)}, " ")
}
