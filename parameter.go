package openweather

import (
	"fmt"
	"reflect"
	"strings"
)

type Parameter struct {
	Latitude  *float32  `json:"lat"`
	Longitude *float32  `json:"lon"`
	AppID     *string   `json:"appid"`
	Exclude   *[]string `json:"exclude"`
	Units     *[]string `json:"units"`
	Lang      *string   `json:"lang"`
}

func (p Parameter) ToQuery() string {
	v := reflect.ValueOf(p)
	var queryParams []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		if value.IsNil() {
			continue
		}

		fieldName := field.Tag.Get("json")
		fieldValue := fmt.Sprintf("%v", value.Elem().Interface())

		// Check if the field value is a slice
		if value.Elem().Kind() == reflect.Slice {
			slice := value.Elem()
			var sliceValues []string
			for j := 0; j < slice.Len(); j++ {
				sliceValue := fmt.Sprintf("%v", slice.Index(j).Interface())
				sliceValues = append(sliceValues, sliceValue)
			}
			fieldValue = strings.Join(sliceValues, ",")
		}

		queryParam := fieldName + "=" + fieldValue
		queryParams = append(queryParams, queryParam)
	}

	return strings.Join(queryParams, "&")
}
