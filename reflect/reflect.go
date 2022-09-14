package reflect

import (
	"reflect"
	"strconv"
	"strings"
)

func CreateQuery(q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	if t.Kind() != reflect.Struct {
		panic("not support type")
	}
	tableName := t.Name()
	var (
		builder       strings.Builder
		columnBuilder strings.Builder
		valueBuilder  strings.Builder
	)
	builder.WriteString("INSERT INTO ")
	builder.WriteString(tableName)

	columnBuilder.WriteString("(")
	valueBuilder.WriteString("VALUES (")
	for i := 0; i < t.NumField(); i++ {
		var val string
		vField := v.Field(i)
		switch v.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(vField.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			val = strconv.FormatUint(vField.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(vField.Float(), 'g', -1, vField.Type().Bits())
		case reflect.Bool:
			val = strconv.FormatBool(vField.Bool())
		case reflect.String:
			val = v.Field(i).String()
		default:
			panic("invalid type")
		}

		if i != 0 {
			columnBuilder.WriteString(", ")
			valueBuilder.WriteString(", ")
		}
		columnBuilder.WriteString(t.Field(i).Name)
		valueBuilder.WriteString(val)
	}
	columnBuilder.WriteString(") ")
	valueBuilder.WriteString(");")

	builder.WriteString(columnBuilder.String())
	builder.WriteString(valueBuilder.String())

	return builder.String()
}
