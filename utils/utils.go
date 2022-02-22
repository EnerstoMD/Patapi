package utils

import (
	"reflect"
	"strings"
)

func PrepareSQLInsertStatement(T interface{}) (string, error) {
	t := reflect.TypeOf(T)
	rv := reflect.ValueOf(T)
	var args, vals []string
	for i := 0; i < t.NumField(); i++ {
		valueptr := rv.Field(i)
		if valueptr.Kind() == reflect.Ptr {
			valueptr = valueptr.Elem()
		}

		value := `'` + valueptr.String() + `'`
		arg, filled := t.Field(i).Tag.Lookup("sql")
		if !filled || !valueptr.IsValid() {
			continue
		}
		args = append(args, arg)
		vals = append(vals, value)
	}
	queryargs := `(` + strings.Join(args, ",") + `)`
	queryvals := `(` + strings.Join(vals, ",") + `)`
	query := `INSERT INTO ` + strings.ToLower(t.Name()) + ` ` + queryargs + ` VALUES ` + queryvals
	return query, nil
}
