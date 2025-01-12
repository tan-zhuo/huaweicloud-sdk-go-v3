package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateLoadbalancerTagsRequestBody This is a auto create Body Object
type BatchCreateLoadbalancerTagsRequestBody struct {

	// 操作类型。 取值范围：create - 创建标签。
	Action BatchCreateLoadbalancerTagsRequestBodyAction `json:"action"`

	// 标签对象列表。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsRequestBody", string(data)}, " ")
}

type BatchCreateLoadbalancerTagsRequestBodyAction struct {
	value string
}

type BatchCreateLoadbalancerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateLoadbalancerTagsRequestBodyAction
}

func GetBatchCreateLoadbalancerTagsRequestBodyActionEnum() BatchCreateLoadbalancerTagsRequestBodyActionEnum {
	return BatchCreateLoadbalancerTagsRequestBodyActionEnum{
		CREATE: BatchCreateLoadbalancerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateLoadbalancerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateLoadbalancerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateLoadbalancerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
