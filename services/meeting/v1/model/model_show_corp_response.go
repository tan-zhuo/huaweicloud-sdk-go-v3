package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCorpResponse Response Object
type ShowCorpResponse struct {
	BasicInfo *QueryCorpBasicResultDto `json:"basicInfo,omitempty"`

	AdminInfo *QueryAdminResultDto `json:"adminInfo,omitempty"`

	ResInfo *QueryCorpResResultDto `json:"resInfo,omitempty"`

	GroupDTO *OrgGroupDto `json:"groupDTO,omitempty"`

	// 企业id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCorpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpResponse", string(data)}, " ")
}
