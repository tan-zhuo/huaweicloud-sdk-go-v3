package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type ImportDataServiceExcelRequestBody struct {

	// 导入包含api的excel文件，大小必须小于4MB。
	File *def.FilePart `json:"file"`
}

func (o ImportDataServiceExcelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataServiceExcelRequestBody struct{}"
	}

	return strings.Join([]string{"ImportDataServiceExcelRequestBody", string(data)}, " ")
}

func (o *ImportDataServiceExcelRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
}
