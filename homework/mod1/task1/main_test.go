package main

import (
	"reflect"
	"testing"
)

func TestShiftSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty slice",
			args: args{s: []int{}},
			want: []int{},
		},
		{
			name: "Slice with one element",
			args: args{s: []int{12}},
			want: []int{12},
		},
		{
			name: "Slice with two elements",
			args: args{s: []int{12, 22}},
			want: []int{22, 12},
		},
		{
			name: "Slice with several elements",
			args: args{s: []int{12, 22, 33, 44}},
			want: []int{22, 33, 44, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiftSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShiftSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty slice",
			args: args{s: []int{}},
			want: []int{},
		},
		{
			name: "Slice with one element",
			args: args{s: []int{12}},
			want: []int{12},
		},
		{
			name: "Slice with two elements",
			args: args{s: []int{12, 22}},
			want: []int{22, 12},
		},
		{
			name: "Slice with even length",
			args: args{s: []int{12, 22, 33, 44}},
			want: []int{44, 33, 22, 12},
		},
		{
			name: "Slice with odd length",
			args: args{s: []int{12, 22, 33, 44, 55}},
			want: []int{55, 44, 33, 22, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}