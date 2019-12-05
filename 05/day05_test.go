package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_parseInstruction(t *testing.T) {
	type args struct {
		instV int
	}
	tests := []struct {
		args  args
		want  int
		want1 []parameterMode
	}{
		{args{1102}, 2, []parameterMode{1, 1, 0}},
		{args{11002}, 2, []parameterMode{0, 1, 1}},
		{args{101}, 1, []parameterMode{1, 0, 0}},
		{args{3}, 3, []parameterMode{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args), func(t *testing.T) {
			got, got1 := parseInstruction(tt.args.instV)
			if got != tt.want {
				t.Errorf("parseInstruction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInstruction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
