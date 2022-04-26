package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func PrepareSQLInsertStatement(c *gin.Context, T interface{}) (string, error) {
	t := reflect.TypeOf(T)
	rv := reflect.ValueOf(T)
	var args, vals []string
	for i := 0; i < t.NumField(); i++ {
		valueptr := rv.Field(i)
		if valueptr.Kind() == reflect.Ptr {
			valueptr = valueptr.Elem()
		}
		value := `'` + valueptr.String() + `'`
		arg, filled := t.Field(i).Tag.Lookup("db")
		if !filled || !valueptr.IsValid() {
			continue
		}
		args = append(args, arg)
		vals = append(vals, value)
	}
	table := strings.ToLower(t.Name())
	if table == "event" {
		queryargs := `(` + strings.Join(args, ",") + `,created_by)`
		queryvals := `(` + strings.Join(vals, ",") + `,` + fmt.Sprintf("%v", c.Keys["userId"]) + `)`
		query := `INSERT INTO public.` + strings.ToLower(t.Name()) + ` ` + queryargs + ` VALUES ` + queryvals
		return query, nil
	}
	queryargs := `(` + strings.Join(args, ",") + `)`
	queryvals := `(` + strings.Join(vals, ",") + `)`

	query := `INSERT INTO public.` + table + ` ` + queryargs + ` VALUES ` + queryvals
	return query, nil
}

func PrepareSQLUpdateStatement(T interface{}, id string) (string, error) {
	t := reflect.TypeOf(T)
	rv := reflect.ValueOf(T)
	var setColums []string
	for i := 0; i < t.NumField(); i++ {
		valueptr := rv.Field(i)
		if valueptr.Kind() == reflect.Ptr {
			valueptr = valueptr.Elem()
		}
		value := `'` + valueptr.String() + `'`
		arg, filled := t.Field(i).Tag.Lookup("db")
		if !filled || !valueptr.IsValid() || arg == "id" {
			continue
		}
		colomnEqualVal := arg + `=` + value
		setColums = append(setColums, colomnEqualVal)
	}
	if len(setColums) == 0 {
		return "", errors.New("no attributes to update")
	}
	query := `UPDATE public.` + strings.ToLower(t.Name()) + ` SET ` + strings.Join(setColums, ",") + ` WHERE id= ` + id
	return query, nil
}
