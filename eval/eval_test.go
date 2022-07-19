package eval

import (
	"bytes"
	"math"
	"testing"
	"text/template"
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

func TestTemplate(t *testing.T) {
	strTmpl := `
[
{{- $v1 := index . "scotty_test_num" -}}
{{- $v2 := index . "scotty_test_den" -}}
{{- $length1 := len $v1 -}}
{{- $length2 := len $v2 -}}
{{- if eq $length1 $length2 -}}
	{{- range $idx, $val := $v1 -}}
		{{$d := index $v2 $idx}}
		{{- if eq .date $d.date -}}
			{{- $rate := printf "%v/%v" .data $d.data -}}
			{{- $date := printf "%.f" .date -}}
			{"date": "{{$date}}", "rate": {{eval $rate}}}
			{{- $lastIdx := add $length1 -1 -}}
			{{- if ne $idx $lastIdx -}},{{- end -}}
		{{- end -}}
	{{- end -}}
{{- end -}}
]
`
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"eval": Eval,
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(strTmpl)
	if err != nil {
		t.Fatal(err)
	}
	data := make(map[string][]*map[string]interface{})
	data["scotty_test_num"] = []*map[string]interface{}{
		{
			"date":       float64(20220713),
			"biz_module": "scotty_test_num",
			"data":       2,
		},
		{
			"date":       float64(20220714),
			"biz_module": "scotty_test_num",
			"data":       1,
		},
	}
	data["scotty_test_den"] = []*map[string]interface{}{
		{
			"date":       float64(20220713),
			"biz_module": "scotty_test_den",
			"data":       5,
		},
		{
			"date":       float64(20220714),
			"biz_module": "scotty_test_den",
			"data":       2,
		},
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success, buf: %v\n", buf.String())
}
