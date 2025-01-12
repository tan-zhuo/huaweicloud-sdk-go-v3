package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistorySlot
type HistorySlot struct {

	// 槽位名称。
	SlotName string `json:"slot_name"`

	// 槽信息。
	SlotValues *[]HistorySlotWord `json:"slot_values,omitempty"`

	// 用户设置的槽位标识。
	SlotIdentification string `json:"slot_identification"`
}

func (o HistorySlot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistorySlot struct{}"
	}

	return strings.Join([]string{"HistorySlot", string(data)}, " ")
}
