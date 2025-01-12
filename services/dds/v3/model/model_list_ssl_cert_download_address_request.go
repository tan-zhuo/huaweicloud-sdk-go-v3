package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSslCertDownloadAddressRequest Request Object
type ListSslCertDownloadAddressRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`
}

func (o ListSslCertDownloadAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSslCertDownloadAddressRequest struct{}"
	}

	return strings.Join([]string{"ListSslCertDownloadAddressRequest", string(data)}, " ")
}
