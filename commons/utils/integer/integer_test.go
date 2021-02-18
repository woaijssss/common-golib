package integer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		sliceA []int
		sliceB []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"intersect", args{[]int{1, 2}, []int{2, 3, 4}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.sliceA, tt.args.sliceB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	type args struct {
		num   int
		slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"intersect success", args{2, []int{2, 3, 4}}, true},
		{"intersect fail", args{10, []int{2, 3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.num, tt.args.slice); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	a := Subtract([]int{1, 2, 3}, []int{3, 4})
	fmt.Printf("%v", a)
}

func TestStrToIntSlice(t *testing.T) {
	fmt.Printf("%v\n", ToIntSlice("", ","))
	fmt.Printf("%v\n", ToIntSlice("1,2,", ","))
}
