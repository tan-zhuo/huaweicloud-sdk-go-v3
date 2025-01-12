package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNotificationsRequest Request Object
type ListNotificationsRequest struct {

	// 通知类型。
	NotificationType ListNotificationsRequestNotificationType `json:"notification_type"`

	// 标识关键操作通知名称。 在不传入该字段的情况下，将查询当前租户所有的关键操作通知。
	NotificationName *string `json:"notification_name,omitempty"`
}

func (o ListNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationsRequest", string(data)}, " ")
}

type ListNotificationsRequestNotificationType struct {
	value string
}

type ListNotificationsRequestNotificationTypeEnum struct {
	SMN ListNotificationsRequestNotificationType
	FUN ListNotificationsRequestNotificationType
}

func GetListNotificationsRequestNotificationTypeEnum() ListNotificationsRequestNotificationTypeEnum {
	return ListNotificationsRequestNotificationTypeEnum{
		SMN: ListNotificationsRequestNotificationType{
			value: "smn",
		},
		FUN: ListNotificationsRequestNotificationType{
			value: "fun",
		},
	}
}

func (c ListNotificationsRequestNotificationType) Value() string {
	return c.value
}

func (c ListNotificationsRequestNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNotificationsRequestNotificationType) UnmarshalJSON(b []byte) error {
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
