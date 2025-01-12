package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddCorpDto struct {
	BasicInfo *CorpBasicDto `json:"basicInfo"`

	AdminInfo *AdminDto `json:"adminInfo"`

	ResInfo *AddCorpResDto `json:"resInfo,omitempty"`

	// 媒体接入（包括SBC和MCU）分组id，可通过[[SP管理员查询资源信息](https://support.huaweicloud.com/api-meeting/meeting_21_1537.html)](tag:hws)[[SP管理员查询资源信息](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_1537.html)](tag:hk)接口查询获取。
	GroupId *string `json:"groupId,omitempty"`

	// 可配置项信息。
	PropertyInfo *[]OrgPropertyDto `json:"propertyInfo,omitempty"`
}

func (o AddCorpDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpDto struct{}"
	}

	return strings.Join([]string{"AddCorpDto", string(data)}, " ")
}
