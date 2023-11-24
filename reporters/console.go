package reporters

import (
	"fmt"
	"reflect"
	"reporter/types"
)

type ConsoleReporter struct {
	Config types.Config
}

func Display(typ reflect.Type) {
	fields := reflect.VisibleFields(typ)

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		if field.Type.Kind() == reflect.Struct {
			fmt.Println("---", field.Type.Name(), "---")
			Display(field.Type)
			continue
		}

		value := reflect.ValueOf(field)
		fmt.Println(field.Name, "-", value.Interface())
	}
}

func (c ConsoleReporter) Report(context types.DroneContext) {
	typ := reflect.TypeOf(context)

	Display(typ)
}
