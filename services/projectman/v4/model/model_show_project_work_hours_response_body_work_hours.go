package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowProjectWorkHoursResponseBodyWorkHours struct {

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 工时日期
	WorkDate *string `json:"work_date,omitempty"`

	// 工时花费
	WorkHoursNum *string `json:"work_hours_num,omitempty"`

	// 工时内容
	Summary *string `json:"summary,omitempty"`

	// 工时类型
	WorkHoursTypeName *string `json:"work_hours_type_name,omitempty"`

	// 工作项id
	IssueId *int32 `json:"issue_id,omitempty"`

	// 工作项类型
	IssueType *string `json:"issue_type,omitempty"`

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 工作项创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 工作项结束时间
	ClosedTime *string `json:"closed_time,omitempty"`

	// 工时创建时间
	WorkHoursCreatedTime *string `json:"work_hours_created_time,omitempty"`

	// 工时更新时间
	WorkHoursUpdatedTime *string `json:"work_hours_updated_time,omitempty"`
}

func (o ShowProjectWorkHoursResponseBodyWorkHours) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursResponseBodyWorkHours struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursResponseBodyWorkHours", string(data)}, " ")
}
