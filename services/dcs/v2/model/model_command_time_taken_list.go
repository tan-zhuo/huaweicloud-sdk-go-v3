package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommandTimeTakenList 命令耗时统计列表
type CommandTimeTakenList struct {

	// 执行命令的总次数
	TotalNum int32 `json:"total_num"`

	// 执行命令的总耗时
	TotalUsecSum float64 `json:"total_usec_sum"`

	// 命令耗时统计结果
	Result CommandTimeTakenListResult `json:"result"`

	// 命令耗时统计
	CommandList []CommandTimeTaken `json:"command_list"`
}

func (o CommandTimeTakenList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandTimeTakenList struct{}"
	}

	return strings.Join([]string{"CommandTimeTakenList", string(data)}, " ")
}

type CommandTimeTakenListResult struct {
	value string
}

type CommandTimeTakenListResultEnum struct {
	SUCCEED CommandTimeTakenListResult
	FAILED  CommandTimeTakenListResult
}

func GetCommandTimeTakenListResultEnum() CommandTimeTakenListResultEnum {
	return CommandTimeTakenListResultEnum{
		SUCCEED: CommandTimeTakenListResult{
			value: "succeed",
		},
		FAILED: CommandTimeTakenListResult{
			value: "failed",
		},
	}
}

func (c CommandTimeTakenListResult) Value() string {
	return c.value
}

func (c CommandTimeTakenListResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommandTimeTakenListResult) UnmarshalJSON(b []byte) error {
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
