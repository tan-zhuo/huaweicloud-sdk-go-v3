package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDiagnosisResponse Response Object
type ListInstanceDiagnosisResponse struct {

	// diagnosis info
	Diagnosis      *[]DiagnosisItemResult `json:"diagnosis,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListInstanceDiagnosisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDiagnosisResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceDiagnosisResponse", string(data)}, " ")
}
