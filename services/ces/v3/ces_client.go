package v3

import (
	httpclient "github.com/tan-zhuo/huaweicloud-sdk-go-v3/core"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/services/ces/v3/model"
)

type CesClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCesClient(hcClient *httpclient.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListAgentStatus 插件状态查询
//
// 插件状态查询，包括uniagent状态以及插件状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAgentStatus(request *model.ListAgentStatusRequest) (*model.ListAgentStatusResponse, error) {
	requestDef := GenReqDefForListAgentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentStatusResponse), nil
	}
}

// ListAgentStatusInvoker 插件状态查询
func (c *CesClient) ListAgentStatusInvoker(request *model.ListAgentStatusRequest) *ListAgentStatusInvoker {
	requestDef := GenReqDefForListAgentStatus()
	return &ListAgentStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateAgentInvocations 批量创建Agent任务
//
// 批量创建Agent任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchCreateAgentInvocations(request *model.BatchCreateAgentInvocationsRequest) (*model.BatchCreateAgentInvocationsResponse, error) {
	requestDef := GenReqDefForBatchCreateAgentInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateAgentInvocationsResponse), nil
	}
}

// BatchCreateAgentInvocationsInvoker 批量创建Agent任务
func (c *CesClient) BatchCreateAgentInvocationsInvoker(request *model.BatchCreateAgentInvocationsRequest) *BatchCreateAgentInvocationsInvoker {
	requestDef := GenReqDefForBatchCreateAgentInvocations()
	return &BatchCreateAgentInvocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgentInvocations 查询Agent任务列表
//
// 查询Agent任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAgentInvocations(request *model.ListAgentInvocationsRequest) (*model.ListAgentInvocationsResponse, error) {
	requestDef := GenReqDefForListAgentInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentInvocationsResponse), nil
	}
}

// ListAgentInvocationsInvoker 查询Agent任务列表
func (c *CesClient) ListAgentInvocationsInvoker(request *model.ListAgentInvocationsRequest) *ListAgentInvocationsInvoker {
	requestDef := GenReqDefForListAgentInvocations()
	return &ListAgentInvocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
