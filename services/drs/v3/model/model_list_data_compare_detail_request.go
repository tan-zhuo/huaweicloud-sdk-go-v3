package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDataCompareDetailRequest Request Object
type ListDataCompareDetailRequest struct {

	// 请求语言类型。
	XLanguage *ListDataCompareDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`

	// 对比状态。 - 0：对比不一致 - 2：对比一致 - 3：目标库表不存在 - 4：对比失败 - 5：正在对比中 - 6：等待对比中 - 7：任务已取消 - 8：源库为空 - 9：目标库为空 - 10：源库和目标库都为空 - 11：源表不存在 - 12：目标表不存在 - 13：原表和目标表都不存在 - 14：源数据库连接失败 - 15：目标库数据库连接失败 - 16：源数据库执行SQL超时 - 17：目标数据库执行SQL超时 - 18：源数据库执行SQL错误 - 19：目标数据库执行SQL错误 - 20：源库和目标库都不存在 - 21：源库不存在 - 22：目标库不存在 - 23：行数为亿行，未进行对比 - 27：超时
	Status *int32 `json:"status,omitempty"`

	// 类型。 - compare：查询正常对比的项 - unCompare：查询无法对比的项
	Type *string `json:"type,omitempty"`

	// 源数据库名。
	DbName *string `json:"db_name,omitempty"`

	// 目标数据库名。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 表名过滤关键字。
	QueryTbName *string `json:"query_tb_name,omitempty"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDataCompareDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataCompareDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDataCompareDetailRequest", string(data)}, " ")
}

type ListDataCompareDetailRequestXLanguage struct {
	value string
}

type ListDataCompareDetailRequestXLanguageEnum struct {
	EN_US ListDataCompareDetailRequestXLanguage
	ZH_CN ListDataCompareDetailRequestXLanguage
}

func GetListDataCompareDetailRequestXLanguageEnum() ListDataCompareDetailRequestXLanguageEnum {
	return ListDataCompareDetailRequestXLanguageEnum{
		EN_US: ListDataCompareDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListDataCompareDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListDataCompareDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListDataCompareDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDataCompareDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
