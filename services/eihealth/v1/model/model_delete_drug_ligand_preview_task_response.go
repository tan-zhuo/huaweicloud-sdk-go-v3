package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugLigandPreviewTaskResponse Response Object
type DeleteDrugLigandPreviewTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugLigandPreviewTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugLigandPreviewTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugLigandPreviewTaskResponse", string(data)}, " ")
}
