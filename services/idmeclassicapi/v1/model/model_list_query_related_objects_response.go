package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryRelatedObjectsResponse Response Object
type ListQueryRelatedObjectsResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Data *[]BasicObjectQueryViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryRelatedObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryRelatedObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryRelatedObjectsResponse", string(data)}, " ")
}
