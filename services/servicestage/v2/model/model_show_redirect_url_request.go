package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRedirectUrlRequest Request Object
type ShowRedirectUrlRequest struct {

	// 仓库类型。 取值范围：github、gitlab、gitee、bitbucket。
	RepoType ShowRedirectUrlRequestRepoType `json:"repo_type"`

	// 站点标签。 比如国际站的，?tag=intl。 默认为空。
	Tag *string `json:"tag,omitempty"`
}

func (o ShowRedirectUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedirectUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowRedirectUrlRequest", string(data)}, " ")
}

type ShowRedirectUrlRequestRepoType struct {
	value string
}

type ShowRedirectUrlRequestRepoTypeEnum struct {
	GITHUB    ShowRedirectUrlRequestRepoType
	GITLAB    ShowRedirectUrlRequestRepoType
	GITEE     ShowRedirectUrlRequestRepoType
	BITBUCKET ShowRedirectUrlRequestRepoType
}

func GetShowRedirectUrlRequestRepoTypeEnum() ShowRedirectUrlRequestRepoTypeEnum {
	return ShowRedirectUrlRequestRepoTypeEnum{
		GITHUB: ShowRedirectUrlRequestRepoType{
			value: "github",
		},
		GITLAB: ShowRedirectUrlRequestRepoType{
			value: "gitlab",
		},
		GITEE: ShowRedirectUrlRequestRepoType{
			value: "gitee",
		},
		BITBUCKET: ShowRedirectUrlRequestRepoType{
			value: "bitbucket",
		},
	}
}

func (c ShowRedirectUrlRequestRepoType) Value() string {
	return c.value
}

func (c ShowRedirectUrlRequestRepoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRedirectUrlRequestRepoType) UnmarshalJSON(b []byte) error {
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
