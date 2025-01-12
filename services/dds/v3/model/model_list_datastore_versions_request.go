package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatastoreVersionsRequest Request Object
type ListDatastoreVersionsRequest struct {

	// 数据库版本类型。取值为“DDS-Community”。
	DatastoreName ListDatastoreVersionsRequestDatastoreName `json:"datastore_name"`
}

func (o ListDatastoreVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoreVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoreVersionsRequest", string(data)}, " ")
}

type ListDatastoreVersionsRequestDatastoreName struct {
	value string
}

type ListDatastoreVersionsRequestDatastoreNameEnum struct {
	DDS_COMMUNITY ListDatastoreVersionsRequestDatastoreName
	DDS_ENHANCED  ListDatastoreVersionsRequestDatastoreName
}

func GetListDatastoreVersionsRequestDatastoreNameEnum() ListDatastoreVersionsRequestDatastoreNameEnum {
	return ListDatastoreVersionsRequestDatastoreNameEnum{
		DDS_COMMUNITY: ListDatastoreVersionsRequestDatastoreName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListDatastoreVersionsRequestDatastoreName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListDatastoreVersionsRequestDatastoreName) Value() string {
	return c.value
}

func (c ListDatastoreVersionsRequestDatastoreName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatastoreVersionsRequestDatastoreName) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
