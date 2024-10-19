package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectRelationsRequest Request Object
type CreateDataobjectRelationsRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 关联主体dataobject所属数据类，小写复数，如告警为alerts，事件为incidents
	DataclassType string `json:"dataclass_type"`

	// 关联主体dataobject的id
	DataObjectId string `json:"data_object_id"`

	// 被关联的dataobject所属数据类，小写复数，如告警为alerts，事件为incidents
	RelatedDataclassType string `json:"related_dataclass_type"`

	Body *CreateDataobjectRelationsRequestBody `json:"body,omitempty"`
}

func (o CreateDataobjectRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectRelationsRequest struct{}"
	}

	return strings.Join([]string{"CreateDataobjectRelationsRequest", string(data)}, " ")
}
