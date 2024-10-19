package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageResponse Response Object
type CreateImageResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId string `json:"job_id"`

	// 创建镜像的ID。
	ImageId        *string `json:"image_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageResponse struct{}"
	}

	return strings.Join([]string{"CreateImageResponse", string(data)}, " ")
}
