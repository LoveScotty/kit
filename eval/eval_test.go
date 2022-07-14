package eval

import (
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		out  float64
	}{
		{
			name: "test1",
			args: args{
				input: "-+-(--(---7+++1))+3*(10/(12/(3+1)-1))/(2+3)-5-3+(8.5)", // 等价于-6+3+0.5
			},
			out: -2.5,
		},
		{
			name: "skip space",
			args: args{
				input: "-   - - 1+ 2*   3  /  4   ",
			},
			out: 0.5,
		},
		{
			name: "test3",
			args: args{
				input: "-+-1+2*3/4",
			},
			out: 2.5,
		},
		{
			name: "test4",
			args: args{
				input: "1/3",
			},
			out: 0.3333333333333333,
		},
		{
			name: "test5",
			args: args{
				input: "",
			},
			out: 0,
		},
		{
			name: "test6",
			args: args{
				input: "1/0",
			},
			out: math.Inf(1),
		},
		{
			name: "test7",
			args: args{
				input: "1/-0",
			},
			out: math.Inf(1),
		},
		{
			name: "test7",
			args: args{
				input: "2/10%5%1",
			},
			out: 0.2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.args.input)
			out := l.Run()
			if out != tt.out {
				t.Errorf("have: %v but want: %v", out, tt.out)
			} else {
				t.Logf("success: %v = %v\n", tt.args.input, out)
			}
		})
	}
}
