package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerselfResourceRecordDetailsResponse Response Object
type ListCustomerselfResourceRecordDetailsResponse struct {

	// 资源详单数据记录。 具体请参见表1。
	MonthlyRecords *[]MonthlyBillRes `json:"monthly_records,omitempty"`

	// 结果集数量，只有成功才返回这个参数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 货币单位代码： USD：美元
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCustomerselfResourceRecordDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordDetailsResponse", string(data)}, " ")
}
