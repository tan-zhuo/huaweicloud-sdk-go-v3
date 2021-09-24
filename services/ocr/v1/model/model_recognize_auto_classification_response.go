package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeAutoClassificationResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *[]AutoClassificationResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RecognizeAutoClassificationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeAutoClassificationResponse struct{}"
	}

	return strings.Join([]string{"RecognizeAutoClassificationResponse", string(data)}, " ")
}