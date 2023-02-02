package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 脚本版本的详细信息。
type ScriptVersion struct {

	// 脚本内容，脚本内容不能为空
	Content string `json:"content"`

	// 创建人，比如为：张三
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 脚本名称，支持数字，下划线，大小写字母 ,中文
	Name *string `json:"name,omitempty"`

	// 租户从IAM申请到的projectid，一般为32位字符串
	ProjectId *string `json:"project_id,omitempty"`

	// 脚本引用次数，脚本被任务和模板引用的次数。默认是0次,引用次数未非负整数，不能出现负数
	ReferenceNumber *int32 `json:"reference_number,omitempty"`

	// 版本所在的脚本id,根据project_id、脚本名称的和，用guava计算hash(md5)方式获取，比如为：d648c682ef5750f700c050b41692b2b8
	ScriptId *string `json:"script_id,omitempty"`

	// 脚本语言，目前支持四种，分别是：SHELL BAT PYTHON POWER_SHELL
	ScriptLanguage *string `json:"script_language,omitempty"`

	// 状态说明  0代表 未上线，1代表已上线  2代表已下线   3代表已禁用
	StatusDesc *int32 `json:"status_desc,omitempty"`

	// 修改人
	UpdateBy *string `json:"update_by,omitempty"`

	// 实体的最后更新时间戳。 注意：执行创建/修补/删除操作时，update_time将更新。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 版本id，唯一标识，根据project_id、脚本名称、脚本版本号的和，用guava计算hash(md5)方式获取，比如为：d648c682ef5750f700c050b41692b2b8
	VersionId *string `json:"version_id,omitempty"`

	// 脚本版本号，支持数字，下划线，大小写字母和小数点
	VersionNumber *string `json:"version_number,omitempty"`

	// 脚本引用详情
	ScriptReference *[]ScriptReferenceDetail `json:"script_reference,omitempty"`
}

func (o ScriptVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptVersion struct{}"
	}

	return strings.Join([]string{"ScriptVersion", string(data)}, " ")
}