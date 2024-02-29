package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUpdateAndCheckinDtoVersionModel struct {
	Data *VersionModel `json:"data,omitempty"`

	// 主对象ID。
	MasterId int64 `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionUpdateAndCheckinDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUpdateAndCheckinDtoVersionModel struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUpdateAndCheckinDtoVersionModel", string(data)}, " ")
}
