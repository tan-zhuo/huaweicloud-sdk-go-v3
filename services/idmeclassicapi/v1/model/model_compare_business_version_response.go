package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareBusinessVersionResponse Response Object
type CompareBusinessVersionResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Data *[]interface{} `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CompareBusinessVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareBusinessVersionResponse struct{}"
	}

	return strings.Join([]string{"CompareBusinessVersionResponse", string(data)}, " ")
}
