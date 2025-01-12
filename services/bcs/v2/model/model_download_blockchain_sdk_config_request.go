package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBlockchainSdkConfigRequest Request Object
type DownloadBlockchainSdkConfigRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`

	Body *CfgRequestBody `json:"body,omitempty"`
}

func (o DownloadBlockchainSdkConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBlockchainSdkConfigRequest struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainSdkConfigRequest", string(data)}, " ")
}
