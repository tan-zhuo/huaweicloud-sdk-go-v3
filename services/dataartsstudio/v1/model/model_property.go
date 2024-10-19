package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Property 节点属性。每种节点类型有自己的定义。
type Property struct {

	// 属性名称。 1.SQL类 HiveSQL、SparkSQL、DWSSQL、DLISQL、RDSSQL：   取值如下：     scriptName，脚本名称     database，数据库名称     connectionName，连接名称     scriptArgs，脚本参数，key、value形式，多个参数间用\\n分隔，例如key1=value1\\nkey2=value2 2.Host类 Shell Python   取值如下：     scriptName，脚本名称     connectionName，连接名称     arguments，脚本参数，key、value形式，多个参数间用\\n分隔，例如key1=value1\\nkey2=value2 3.CDMJob   取值如下：     clusterName，集群名称     jobName，作业名称 4.DISTransferTask   取值如下：     streamName，DIS通道名称     destinationType，转储目标     duplicatePolicy，转储任务重名处理策略     configuration，转储配置 5.OBSManager   取值如下：     action，动作类型     path，OBS路径 6.RESTAPI   取值如下：     url，URL地址     method，HTTP方法     headers, HTTP消息头，每个消息头的格式为<消息头名称>=<值>，多个消息头之间使用换行符分割。     body, 消息体 7.SMN   取值如下：     topic，SMN主题URN     subject，消息标题，给邮箱订阅者发送邮件时作为邮件主题。     messageType, 消息类型     message, 发送的消息。 8.MRSSpark、MapReduce   取值如下：     clusterName，集群名称     jobName，作业名称     resourcePath，自定义Jar包OBS资源路径     parameters，  Jar包自定义参数；对于开发的自定义Jar包，可以在此处对参数进行输入替换     input，输入路径     output，输出路径     programParameter，运行程序参数；允许多个key:value，多个参数要用竖线隔开。 9.DLISpark   取值如下：     clusterName，集群名称     jobName，作业名称     resourceType，DLI作业运行资源类型     sparkConfig，Spark作业运行参数     jobClass，主类名称。当应用程序类型为“.jar”时，主类名称不能为空。     jarArgs，主类入口参数     resourcePath，JAR包资源路径 10.MRSFlink   取值如下：     clusterName，集群名称     jobName，作业名称     flinkJobType，flink作业类型；FLink SQL作业或Flink JAR作业。     flinkJobProcessType，flink作业处理模式；批处理模式或流处理模式。     scriptName，脚本名称；Flink SQL关联的SQL脚本。     resourcePath，JAR包资源路径     input，输入路径     output，输出路径     programParameter，运行程序参数；允许多个key:value，多个参数要用竖线隔开。 11.MRS HetuEngine   取值如下：     clusterName，集群名称     jobName，作业名称     statementOrScript，脚本类型；SQL语句或关联SQL脚本     statement，自定义的SQL内容。     scriptName，选择关联的SQL脚本。     Data Warehouse，指定HetuEngine服务所需数据连接。     Schema，使用HetuEngine服务所要访问的数据源schema名称。     Database，使用HetuEngine服务所要访问的数据源database名称。     Queue，使用HetuEngine服务所需资源队列名称。
	Name *string `json:"name,omitempty"`

	// 属性值
	Value *string `json:"value,omitempty"`
}

func (o Property) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Property struct{}"
	}

	return strings.Join([]string{"Property", string(data)}, " ")
}
