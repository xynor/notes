package main

import (
	"fmt"
	"html/template"
	"os"
	"reflect"
	"testing"
)

var temp = `<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta charset="utf-8">
</head>
<body>
<table border="1">
	<tr>
		{{range .Rows}}
        <th>{{.}}</th>
		{{end}}
    </tr>
	{{range .Item}}
	<tr>
		{{range rangeStruct .}}
        <td>{{.}}</td>
		{{end}}
	</tr>
	{{end}}
</table>
</body>
</html>`

type Items struct {
	Rows []string
	Item []Item
}

type Item struct {
	Name   string
	Age    int
	Money  float64
	Height int
}

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

func TestTemplate(t *testing.T) {
	file := "./t.html"
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()
	tpl, err := template.New("t").Funcs(templateFuncs).Parse(temp)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	rows, err := FillRows(new(Item))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	items := Items{
		Rows: rows,
		Item: []Item{{Name: "wang", Age: 11, Money: 11.22, Height: 11},
			{Name: "dsa", Age: 23, Money: 54.65, Height: 34},
			{Name: "cfh", Age: 54, Money: 34422, Height: 3445},
			{Name: "sdf", Age: 66, Money: 46.01},
			{Name: "gjf", Age: 234, Money: 34564, Height: 345}},
	}
	err = tpl.Execute(f, items)
	if err != nil {
		fmt.Println(err)
	}
}
func FillRows(item *Item) (row []string, err error) {
	rt := reflect.TypeOf(item)
	if rt.Kind() != reflect.Ptr {
		return row, fmt.Errorf("need a Ptr")
	}
	rte := rt.Elem()
	if rte.Kind() != reflect.Struct {
		return row, fmt.Errorf("need a Ptr")
	}
	for i := 0; i < rte.NumField(); i++ {
		row = append(row, rte.Field(i).Name)
	}
	return
}
func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}
