package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackups2Response Response Object
type ListBackups2Response struct {

	// 备份总个数。请求失败时，字段为空。
	BackupCount *int32 `json:"backup_count,omitempty"`

	// 当前Project ID下的所有图的备份列表。请求失败时，字段为空。
	BackupList     *[]ListBackupsRespBackupList `json:"backup_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListBackups2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackups2Response struct{}"
	}

	return strings.Join([]string{"ListBackups2Response", string(data)}, " ")
}
