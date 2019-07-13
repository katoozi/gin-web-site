package sqltools

import (
	"fmt"
	"reflect"
)

// GenerateInsertQuery will generate insert query for given data
func GenerateInsertQuery(tblName string, data struct{}) {
	// dataStruct, ok := data.(struct{})
	// if !ok {
	// 	log.Fatalf("Error while cast interface to struct: %v", ok)
	// }

	fmt.Println("Start Generate Insert Query...")

	t := reflect.TypeOf(data)
	fmt.Println(t.FieldByName("ID"))
}
