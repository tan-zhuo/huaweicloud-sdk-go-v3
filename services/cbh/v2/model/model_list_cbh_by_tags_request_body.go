package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCbhByTagsRequestBody TMS查询CBH资源实例列表请求体。
type ListCbhByTagsRequestBody struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源。  此时忽略 “tags”、“tags_any”、“not_tags”、“not_tags_any”字段。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含50个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。  Key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系， key-value结构中value是或的关系。  无tag过滤条件时返回全量数据。
	Tags *[]Tags `json:"tags,omitempty"`

	// 包含任意标签，最多包含50个key，每个key下面的value最多10个,每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。  结果返回包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。
	TagsAny *[]Tags `json:"tags_any,omitempty"`

	// 不包含标签，最多包含50个key，每个key下面的value最多10个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。  结果返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。
	NotTags *[]Tags `json:"not_tags,omitempty"`

	// 不包含任意标签，最多包含50个key，每个key下面的value最多10个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。  结果返回不包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。
	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`

	// 仅op_service权限可以使用此字段做资源实例过滤条件。  目前TMS调用时只包含一个tag结构体。  key：_sys_enterprise_project_id  value：企业项目id列表  目前TMS调用时，key下面只包含一个value。0表示默认企业项目  sys_tags和租户标签过滤条件（without_any_tag 、tags、tags_any、not_tags、not_tags_any）不能同时使用  无sys_tags时按照tag接口处理，无tag过滤条件时返回全量数据。。
	SysTags *[]Tags `json:"sys_tags,omitempty"`

	// 搜索字段,key为要匹配的字段，如resource_name等。  value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。  根据key的值确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写），如果value为空字符串精确匹配（多数服务不存在资源名称为空的情况，因此此类情况返回空列表）。  resource_id为精确匹配。第一期只做resource_name，后续再扩展。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListCbhByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCbhByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListCbhByTagsRequestBody", string(data)}, " ")
}
