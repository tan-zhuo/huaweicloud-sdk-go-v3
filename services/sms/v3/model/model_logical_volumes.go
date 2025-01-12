package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalVolumes 逻辑卷信息
type LogicalVolumes struct {

	// 块数量
	BlockCount *int32 `json:"block_count,omitempty"`

	// 块大小
	BlockSize *int64 `json:"block_size,omitempty"`

	// 文件系统
	FileSystem string `json:"file_system"`

	// inode数量
	InodeSize int32 `json:"inode_size"`

	// inode节点数量
	InodeNums *int64 `json:"inode_nums,omitempty"`

	// 分区类型，普通分区，启动分区，系统分区
	DeviceUse *string `json:"device_use,omitempty"`

	// 挂载点
	MountPoint string `json:"mount_point"`

	// 名称
	Name string `json:"name"`

	// 大小
	Size int64 `json:"size"`

	// 使用大小
	UsedSize int64 `json:"used_size"`

	// 剩余空间
	FreeSize int64 `json:"free_size"`
}

func (o LogicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalVolumes struct{}"
	}

	return strings.Join([]string{"LogicalVolumes", string(data)}, " ")
}
