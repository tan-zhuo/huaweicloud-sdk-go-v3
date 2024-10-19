package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Dpd struct {

	// 对等体存活检测空闲时间
	Interval *int32 `json:"interval,omitempty"`

	// 对等体存活检测报文重传间隔
	Timeout *int32 `json:"timeout,omitempty"`

	// 对等体存活检测报文格式
	Msg *DpdMsg `json:"msg,omitempty"`
}

func (o Dpd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dpd struct{}"
	}

	return strings.Join([]string{"Dpd", string(data)}, " ")
}

type DpdMsg struct {
	value string
}

type DpdMsgEnum struct {
	SEQ_HASH_NOTIFY DpdMsg
	SEQ_NOTIFY_HASH DpdMsg
}

func GetDpdMsgEnum() DpdMsgEnum {
	return DpdMsgEnum{
		SEQ_HASH_NOTIFY: DpdMsg{
			value: "seq-hash-notify",
		},
		SEQ_NOTIFY_HASH: DpdMsg{
			value: "seq-notify-hash",
		},
	}
}

func (c DpdMsg) Value() string {
	return c.value
}

func (c DpdMsg) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DpdMsg) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
