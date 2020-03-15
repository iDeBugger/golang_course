package main

import (
	"reflect"
	"testing"
)

func TestShiftSlice(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty slice",
			args: args{s: []int{}, n: 1},
			want: []int{},
		},
		{
			name: "Shift for 1 on slice with one element",
			args: args{s: []int{12}, n: 1},
			want: []int{12},
		},
		{
			name: "Shift for 1 on slice with two elements",
			args: args{s: []int{12, 22}, n: 1},
			want: []int{22, 12},
		},
		{
			name: "Shift for 1 on slice with several elements",
			args: args{s: []int{12, 22, 33, 44}, n: 1},
			want: []int{22, 33, 44, 12},
		},
		{
			name: "Shift for 0 on empty slice",
			args: args{s: []int{}, n: 0},
			want: []int{},
		},
		{
			name: "Shift for 0 on non-empty slice",
			args: args{s: []int{1, 2, 3, 4}, n: 0},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Shift for 2 on empty slice",
			args: args{s: []int{}, n: 2},
			want: []int{},
		},
		{
			name: "Shift for 2 on slice with one element",
			args: args{s: []int{12}, n: 2},
			want: []int{12},
		},
		{
			name: "Shift for 2 on slice with two elements",
			args: args{s: []int{12, 22}, n: 2},
			want: []int{12, 22},
		},
		{
			name: "Shift for 2 on slice with several elements",
			args: args{s: []int{12, 22, 33, 44}, n: 2},
			want: []int{33, 44, 12, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiftSlice(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
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
