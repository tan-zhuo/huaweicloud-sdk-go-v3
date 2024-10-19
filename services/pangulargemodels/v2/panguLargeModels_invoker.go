package v2

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/services/pangulargemodels/v2/model"
)

type ExecuteChatCompletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteChatCompletionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteChatCompletionInvoker) Invoke() (*model.ExecuteChatCompletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteChatCompletionResponse), nil
	}
}

type ExecuteTextCompletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTextCompletionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteTextCompletionInvoker) Invoke() (*model.ExecuteTextCompletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTextCompletionResponse), nil
	}
}
