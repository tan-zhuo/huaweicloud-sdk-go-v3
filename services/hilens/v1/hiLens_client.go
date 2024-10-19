package v1

import (
	httpclient "github.com/tan-zhuo/huaweicloud-sdk-go-v3/core"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/services/hilens/v1/model"
)

type HiLensClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewHiLensClient(hcClient *httpclient.HcHttpClient) *HiLensClient {
	return &HiLensClient{HcClient: hcClient}
}

func HiLensClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListDeviceAlarms 获取设备告警列表
//
// 获取设备告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListDeviceAlarms(request *model.ListDeviceAlarmsRequest) (*model.ListDeviceAlarmsResponse, error) {
	requestDef := GenReqDefForListDeviceAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceAlarmsResponse), nil
	}
}

// ListDeviceAlarmsInvoker 获取设备告警列表
func (c *HiLensClient) ListDeviceAlarmsInvoker(request *model.ListDeviceAlarmsRequest) *ListDeviceAlarmsInvoker {
	requestDef := GenReqDefForListDeviceAlarms()
	return &ListDeviceAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevices 获取基础版设备列表
//
// 获取基础版设备列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}

// ListDevicesInvoker 获取基础版设备列表
func (c *HiLensClient) ListDevicesInvoker(request *model.ListDevicesRequest) *ListDevicesInvoker {
	requestDef := GenReqDefForListDevices()
	return &ListDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
