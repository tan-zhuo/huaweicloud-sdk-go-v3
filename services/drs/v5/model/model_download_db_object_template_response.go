package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// DownloadDbObjectTemplateResponse Response Object
type DownloadDbObjectTemplateResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadDbObjectTemplateResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadDbObjectTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDbObjectTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadDbObjectTemplateResponse", string(data)}, " ")
}
