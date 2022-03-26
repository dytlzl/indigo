package printer

import (
	"fmt"
	"strings"
)

func PrintTable[T any](columnNames []string, data []T, fn func(T) []string, prefix string) {
	table := make([][]string, len(data))

	columnWidthList := make([]int, len(columnNames))
	for x := range columnWidthList {
		columnWidthList[x] = len(columnNames[x])
	}
	for index, element := range data {
		table[index] = fn(element)
		for x := range columnWidthList {
			if columnWidthList[x] < len(table[index][x]) {
				columnWidthList[x] = len(table[index][x])
			}
		}
	}

	for x, cell := range columnNames {
		fmt.Printf("%s%s%s", prefix, cell, strings.Repeat(" ", columnWidthList[x]+3-len(cell)))
	}
	fmt.Println()
	for _, row := range table {
		for x, cell := range row {
			fmt.Printf("%s%s%s", prefix, cell, strings.Repeat(" ", columnWidthList[x]+3-len(cell)))
		}
		fmt.Println()
	}
}
