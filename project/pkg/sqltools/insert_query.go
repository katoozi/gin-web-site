package sqltools

import (
	"fmt"
	"reflect"
	"strings"
)

var insertQuerySchema = `INSERT INTO "%s" (%s) VALUES (%s);`

// GenerateInsertQuery will generate insert query for given data
func GenerateInsertQuery(tblName string, data interface{}) string {
	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.TypeOf(data)
		values := reflect.ValueOf(data)

		fieldsSection := ""
		valuesSection := ""

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			sqlToolsTag := field.Tag.Get("sqltools")
			if sqlToolsTag == "" {
				sqlToolsTag = field.Name
			}
			fieldsSection += sqlToolsTag + ","
			fieldType := values.Field(i).Kind()
			if fieldType == reflect.Int {
				valuesSection += fmt.Sprintf(`'%d',`, values.Field(i).Int())
			} else if fieldType == reflect.String {
				valuesSection += fmt.Sprintf(`'%s',`, values.Field(i))
			}
		}
		fieldsSection = strings.TrimRight(fieldsSection, ",")
		valuesSection = strings.TrimRight(valuesSection, ",")

		return fmt.Sprintf(insertQuerySchema, tblName, fieldsSection, valuesSection)

	}
	return ""
}
