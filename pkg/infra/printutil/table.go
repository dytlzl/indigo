package printutil

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/duration"
)

func PrintTable[T any](data []T) {
	PrintTableWithPrefix(data, "")
}

func PrintTableWithPrefix[T any](data []T, prefix string) {
	table := make([][]string, len(data))
	columns := columnsFromStruct(data)
	columnWidthList := make([]int, len(columns))
	for x := range columnWidthList {
		columnWidthList[x] = len(columns[x].name)
	}
	for index, element := range data {
		table[index] = stringSliceFromStruct(element, columns)
		for x := range columnWidthList {
			if columnWidthList[x] < len(table[index][x]) {
				columnWidthList[x] = len(table[index][x])
			}
		}
	}
	for x, column := range columns {
		if x == len(columns)-1 {
			fmt.Printf("%s%s", prefix, column.name)
		} else {
			fmt.Printf("%s%s%s", prefix, column.name, strings.Repeat(" ", columnWidthList[x]+3-len(column.name)))
		}
	}
	fmt.Println()
	for _, row := range table {
		for x, cell := range row {
			if x == len(columns)-1 {
				fmt.Printf("%s%s", prefix, cell)
			} else {
				fmt.Printf("%s%s%s", prefix, cell, strings.Repeat(" ", columnWidthList[x]+3-len(cell)))
			}
		}
		fmt.Println()
	}
}

func stringSliceFromStruct[T any](strct T, columns []column) []string {
	v := reflect.ValueOf(strct)
	t := v.Type()
	s := make([]string, 0, t.NumField())
	for _, column := range columns {
		s = append(s, formatInterface(v.FieldByName(column.field).Interface()))
	}
	return s
}

func formatInterface(i any) string {
	switch typed := i.(type) {
	case string:
		return typed
	case int:
		return strconv.Itoa(typed)
	case time.Time:
		return duration.HumanDuration(time.Since(typed))
	}
	return ""
}

func columnsFromStruct(slice any) []column {
	v := reflect.ValueOf(slice)
	t := v.Type().Elem()
	columns := make([]column, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		tag := strings.Split(t.Field(i).Tag.Get("print"), ",")
		if len(tag) >= 1 {
			index := 0
			if len(tag) >= 2 {
				index, _ = strconv.Atoi(tag[1])
			}
			columns = append(columns, column{
				index: index,
				name:  tag[0],
				field: t.Field(i).Name,
			})
		}
	}
	sort.Slice(columns, func(i, j int) bool { return columns[i].index < columns[j].index })
	return columns
}

type column struct {
	index int
	name  string
	field string
}
