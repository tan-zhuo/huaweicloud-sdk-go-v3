package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReleaseHistoryRequest Request Object
type ShowReleaseHistoryRequest struct {

	// 模板实例名称
	Name string `json:"name"`

	// 模板实例所在的命名空间
	Namespace string `json:"namespace"`

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ShowReleaseHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReleaseHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowReleaseHistoryRequest", string(data)}, " ")
}
