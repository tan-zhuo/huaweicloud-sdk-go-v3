package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatsByTagsRequestBody 查询资源实例的请求体
type ListNatsByTagsRequestBody struct {

	// 包含标签对象列表，最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]PublicTag `json:"tags,omitempty"`

	// 包含任意标签对象列表，最多包含10个key，每个key下面的value最多10个,结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	TagsAny *[]PublicTag `json:"tags_any,omitempty"`

	// 不包含标签对象列表，最多包含10个key，每个key下面的value最多10个, 结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTags *[]PublicTag `json:"not_tags,omitempty"`

	// 不包含任意标签对象列表，最多包含10个key，每个key下面的value最多10个,结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回不包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTagsAny *[]PublicTag `json:"not_tags_any,omitempty"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *string `json:"limit,omitempty"`

	// （索引位置）， 从offset指定的下一条数据开始查询。查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数（action为count时无此参数）如果action为filter默认为0，必须为数字，不能为负数
	Offset *string `json:"offset,omitempty"`

	// - 操作标识（仅限于filter，count）：filter（过滤），count（查询总条数） - 如果是filter就是分页查询，如果是count只需按照条件将总条数返回即可。
	Action string `json:"action"`

	// - 搜索字段列表，key为要匹配的字段，如resource_name等。value为匹配的值。此字段为固定字典值。 - 根据不同的字段确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写），如果value为空字符串精确匹配。resource_id为精确匹配。
	Matches *[]PublicMatch `json:"matches,omitempty"`
}

func (o ListNatsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListNatsByTagsRequestBody", string(data)}, " ")
}
