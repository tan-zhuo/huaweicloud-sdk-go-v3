package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdministrationAgencyUrnPrimitiveTypeHolder struct {

	// 管理委托URN  资源编排服务使用该委托获取成员账号委托给管理账号的权限。该委托中必须含有sts:tokens:assume权限，用以后续获取被管理委托凭证。如果不包含，则会在新增或者部署实例时报错。  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予administration_agency_urn，administration_agency_name只支持接收委托名称，如果给予了信任委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。
	AdministrationAgencyUrn *string `json:"administration_agency_urn,omitempty"`
}

func (o AdministrationAgencyUrnPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdministrationAgencyUrnPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"AdministrationAgencyUrnPrimitiveTypeHolder", string(data)}, " ")
}
