package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreDatabasesRequest Request Object
type ListRestoreDatabasesRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 索引位置偏移量。   - 索引位置偏移量，表示从指定project ID下最新的专属资源创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的专属资源信息。   - 取值大于或等于0。   - 不传该参数时，查询偏移量默认为0，表示从最新的创建时间对应的专属资源开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询专属资源个数上限值。   - 取值范围：1~100。不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestoreDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreDatabasesRequest", string(data)}, " ")
}
