package namelist

import (
	"reflect"
	"testing"
)

func Test_increaseMap(t *testing.T) {
	type args struct {
		m map[string]int
		s []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"Increasing empty map",
			args{
				map[string]int{},
				[]string{"Gustavo", "Olívia", "Inês", "Eva"},
			},
			map[string]int{
				"Gustavo": 1,
				"Olívia":  1,
				"Inês":    1,
				"Eva":     1,
			},
		},
		{
			"Increasing non-empty map",
			args{
				map[string]int{
					"Gustavo":   1,
					"Sara":      1,
					"Carlos":    1,
					"Magdalena": 1,
				},
				[]string{"Gustavo", "Olívia", "Inês", "Eva", "Magdalena"},
			},
			map[string]int{
				"Gustavo":   2,
				"Olívia":    1,
				"Inês":      1,
				"Eva":       1,
				"Sara":      1,
				"Carlos":    1,
				"Magdalena": 2,
			},
		},
		{
			"Increasing nothing",
			args{
				map[string]int{},
				[]string{},
			},
			map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increaseMap(tt.args.m, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increaseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test 1",
			args{
				[]string{"Gustavo", "Sara", "Carlos", "Magdalena"},
				[]string{"Gustavo", "Olívia", "Inês", "Eva"},
			},
			[]string{"Carlos", "Eva", "Gustavo", "Inês", "Magdalena", "Olívia", "Sara"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newStringSlice(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Map to Slice Conversion",
			args{
				map[string]int{
					"Gustavo":   1,
					"Sara":      1,
					"Carlos":    1,
					"Magdalena": 1,
				},
			},
			[]string{"Gustavo", "Sara", "Carlos", "Magdalena"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStringSlice(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRepeated(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"Count amount",
			args{
				[]string{"Gustavo", "Sara", "Carlos", "Magdalena"},
				[]string{"Gustavo", "Olívia", "Inês", "Eva", "Magdalena"},
			},
			map[string]int{
				"Gustavo":   2,
				"Olívia":    1,
				"Inês":      1,
				"Eva":       1,
				"Sara":      1,
				"Carlos":    1,
				"Magdalena": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRepeated(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRepeated() = %v, want %v", got, tt.want)
			}
		})
	}
}
