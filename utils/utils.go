package utils

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

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
	table := CamelToSnakeCase(t.Name())
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
	query := `UPDATE public.` + CamelToSnakeCase(t.Name()) + ` SET ` + strings.Join(setColums, ",") + ` WHERE id= ` + id
	return query, nil
}

func ReadStructToBeInserted(c *gin.Context, T interface{}) ([]string, error) {
	t := reflect.TypeOf(T)
	rv := reflect.ValueOf(T)
	var args []string
	//var vals []string
	for i := 0; i < t.NumField(); i++ {
		valueptr := rv.Field(i)
		if valueptr.Kind() == reflect.Ptr {
			valueptr = valueptr.Elem()
		}
		//value := `'` + valueptr.String() + `'`
		arg, filled := t.Field(i).Tag.Lookup("db")
		if !filled || !valueptr.IsValid() {
			continue
		}
		args = append(args, arg)
		//vals = append(vals, value)
	}
	table := CamelToSnakeCase(t.Name())
	queryargs := `(` + strings.Join(args, ",") + `)`
	queryvals := `(:` + strings.Join(args, ", :") + `)`
	return []string{table, queryargs, queryvals}, nil
}

func CamelToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
