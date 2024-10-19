package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRuleMetricsResponse Response Object
type ListAlertRuleMetricsResponse struct {

	// 指标类型，分组数量。Metric category. GROUP_COUNT.
	Category *ListAlertRuleMetricsResponseCategory `json:"category,omitempty"`

	// 指标值。Metric value.
	Metric map[string]float32 `json:"metric,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertRuleMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleMetricsResponse", string(data)}, " ")
}

type ListAlertRuleMetricsResponseCategory struct {
	value string
}

type ListAlertRuleMetricsResponseCategoryEnum struct {
	GROUP_COUNT ListAlertRuleMetricsResponseCategory
}

func GetListAlertRuleMetricsResponseCategoryEnum() ListAlertRuleMetricsResponseCategoryEnum {
	return ListAlertRuleMetricsResponseCategoryEnum{
		GROUP_COUNT: ListAlertRuleMetricsResponseCategory{
			value: "GROUP_COUNT",
		},
	}
}

func (c ListAlertRuleMetricsResponseCategory) Value() string {
	return c.value
}

func (c ListAlertRuleMetricsResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRuleMetricsResponseCategory) UnmarshalJSON(b []byte) error {
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
