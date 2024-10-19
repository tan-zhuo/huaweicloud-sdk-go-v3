package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteArtifactFileRequest Request Object
type DeleteArtifactFileRequest struct {
	Body *PathMap `json:"body,omitempty"`
}

func (o DeleteArtifactFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteArtifactFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteArtifactFileRequest", string(data)}, " ")
}
