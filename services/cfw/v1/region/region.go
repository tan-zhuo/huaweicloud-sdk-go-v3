package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_4 = region.NewRegion("cn-south-4", "https://cfw.cn-south-4.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://cfw.cn-southwest-2.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://cfw.cn-south-1.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://cfw.cn-east-3.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://cfw.cn-north-4.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://cfw.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://cfw.ap-southeast-3.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://cfw.la-north-2.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://cfw.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_4 = region.NewRegion("ap-southeast-4", "https://cfw.ap-southeast-4.myhuaweicloud.com")
var LA_SOUTH_2 = region.NewRegion("la-south-2", "https://cfw.la-south-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-4":     CN_SOUTH_4,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-north-4":     CN_NORTH_4,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"la-north-2":     LA_NORTH_2,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-4": AP_SOUTHEAST_4,
	"la-south-2":     LA_SOUTH_2,
}

var provider = region.DefaultProviderChain("CFW")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}