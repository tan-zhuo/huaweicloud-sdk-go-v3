package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceConfigurationResponse Response Object
type ShowInstanceConfigurationResponse struct {

	// 数据库版本名称。
	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	// 数据库名称。
	DatastoreName *string `json:"datastore_name,omitempty"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated *string `json:"updated,omitempty"`

	// 参数模板ID。
	Id *string `json:"id,omitempty"`

	// 数据库实例类型。 GeminiDB Cassandra集群类型为\"Cluster\"。 GeminiDB Mongo副本集类型为\"ReplicaSet\"。 GeminiDB Mongo集群类型为\"Sharding\"。 GeminiDB Influx集群类型为\"Cluster\"。 GeminiDB Influx单节点类型为\"InfluxdbSingle\"。 GeminiDB Redis集群类型为“Cluster”。 GeminiDB Redis主备类型为“Replication”。
	Mode *string `json:"mode,omitempty"`

	// 参数对象，用户基于默认参数模板自定义的参数配置。
	ConfigurationParameters *[]ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                             `json:"-"`
}

func (o ShowInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationResponse", string(data)}, " ")
}
