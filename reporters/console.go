package reporters

import (
	"fmt"
	"reflect"
	"reporter/types"
)

type ConsoleReporter struct {
	Config types.Config
}

func (c ConsoleReporter) Report(context types.DroneContext) {
	typ := reflect.TypeOf(context)
	fields := reflect.VisibleFields(typ)

	for i := 0; i < len(fields); i++ {
		field := fields[i]
		value := reflect.ValueOf(field)
		fmt.Println(field.Name, "-", value.String())
	}
}
