package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsModifierDto struct {

	// ID列表。
	Ids []int64 `json:"ids"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsModifierDto", string(data)}, " ")
}
