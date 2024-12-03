package nats_contracts

import "reflect"

func ConvertDocumentToPlaceholderMap(d *Document) map[string]interface{} {
	r := make(map[string]interface{})

	f := reflect.VisibleFields(reflect.TypeOf(Document{}))
	for _, item := range f {
		r[item.Name] = reflect.Indirect(reflect.ValueOf(d)).FieldByName(item.Name)
	}

	return r
}
