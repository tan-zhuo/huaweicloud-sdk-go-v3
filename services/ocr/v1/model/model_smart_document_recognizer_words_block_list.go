package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerWordsBlockList struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 文字块识别结果的置信度。
	Confidence *float32 `json:"confidence,omitempty"`
}

func (o SmartDocumentRecognizerWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerWordsBlockList struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerWordsBlockList", string(data)}, " ")
}
