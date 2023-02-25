package printutil

import (
	"reflect"
	"testing"
	"time"
)

func Test_stringSliceFromStruct(t *testing.T) {
	tests := []struct {
		name    string
		strct   any
		columns []column
		want    []string
	}{
		{
			name: "return correct value",
			strct: struct {
				ID        int       `print:"id"`
				Name      string    `print:"name"`
				UpdatedAt time.Time `print:"age"`
			}{
				ID:        1,
				Name:      "name01",
				UpdatedAt: time.Now(),
			},
			columns: columnsFromStruct([]struct {
				ID        int       `print:"id"`
				Name      string    `print:"name"`
				UpdatedAt time.Time `print:"age"`
			}{}),
			want: []string{"1", "name01", "0s"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSliceFromStruct(tt.strct, tt.columns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringSliceFromStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_columnNamesFromStruct(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want []column
	}{
		{
			name: "return correct value",
			arg: []struct {
				ID        int       `print:"id"`
				Name      string    `print:"name,1"`
				UpdatedAt time.Time `print:"age,2"`
			}{},
			want: []column{{0, "id", "ID"}, {1, "name", "Name"}, {2, "age", "UpdatedAt"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := columnsFromStruct(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("columnNamesFromStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
