package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadImagesResponse Response Object
type DownloadImagesResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadImagesResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadImagesResponse struct{}"
	}

	return strings.Join([]string{"DownloadImagesResponse", string(data)}, " ")
}
