package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLineGroupResponse Response Object
type DeleteLineGroupResponse struct {

	// 线路分组名称。
	Name *string `json:"name,omitempty"`

	// 线路分组包含的线路列表。 解析线路ID。
	Lines *[]string `json:"lines,omitempty"`

	// 资源状态。 取值范围：PENDING_CREATE，ACTIVE，PENDING_DELETE，PENDING_UPDATE，ERROR，FREEZE，DISABLE。
	Status *string `json:"status,omitempty"`

	// 线路分组的描述信息
	Description *string `json:"description,omitempty"`

	// 线路分组的id。
	LineId *string `json:"line_id,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLineGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteLineGroupResponse", string(data)}, " ")
}
