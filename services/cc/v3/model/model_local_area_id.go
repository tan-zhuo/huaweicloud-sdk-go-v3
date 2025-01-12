package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalAreaId 本端大区ID。
type LocalAreaId struct {
	LocalAreaId *AreaIdDef `json:"local_area_id"`
}

func (o LocalAreaId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalAreaId struct{}"
	}

	return strings.Join([]string{"LocalAreaId", string(data)}, " ")
}
