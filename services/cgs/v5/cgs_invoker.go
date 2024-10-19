package v5

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/services/cgs/v5/model"
)

type ListContainerNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerNodesInvoker) Invoke() (*model.ListContainerNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerNodesResponse), nil
	}
}
