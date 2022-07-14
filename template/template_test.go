package template

import (
	"bytes"
	"testing"
	"text/template"
)

func TestText(t *testing.T) {
	type args struct {
		t    *template.Template
		text string
	}
	tmpl := template.New("test")
	tests := []struct {
		name string
		args args
		out  string
	}{
		{
			name: "test1",
			args: args{
				t:    tmpl,
				text: `{{define "hello name? age?"}}hello {{.name}}, you're {{.age}}{{end}}{{hello "scotty" "18"}}`,
			},
			out: "hello scotty, you're 18",
		},
		{
			name: "test2",
			args: args{
				t:    tmpl,
				text: `{{define "hello name..."}}hello {{.name}}{{end}}{{hello "scotty"}}`,
			},
			out: "hello [scotty]",
		},
		{
			name: "test3",
			args: args{
				t:    tmpl,
				text: `{{define "hello name"}}hello {{.name}}{{end}}{{hello "scotty"}}`,
			},
			out: "hello scotty",
		},
		{
			name: "test4",
			args: args{
				t:    tmpl,
				text: `{{define "hello name?"}}hello {{.name}}{{end}}{{hello}}`,
			},
			out: "hello <no value>", // 没有输入参数会被替换成<no value>
		},
		{
			name: "test5",
			args: args{
				t:    tmpl,
				text: `{{define "hello name? age?"}}hello {{.name}}, you're {{.age}}{{end}}{{hello "scotty" "18"}}`,
			},
			out: "hello scotty, you're 18",
		},
		{
			name: "test6",
			args: args{
				t:    tmpl,
				text: `{{define "hit_rate a? b?"}}{{.a}}*{{.a}} / {{.b}}{{end}}{{hit_rate "2" "4"}}`,
			},
			out: "2*2 / 4",
		},
		{
			name: "test6",
			args: args{
				t:    tmpl,
				text: "{{define \"hit_rate name? a? b?\"}}{\"name\":{{.name}}, \"rate\": {{.a}}/{{.b}}}{{end}}{{hit_rate \"scotty\" 10, 20}}",
			},
			out: "2*2 / 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Parse(tt.args.t, tt.args.text)
			var out string
			if err != nil {
				out = "Parse: " + err.Error()
			} else {
				var buf bytes.Buffer
				err := tmpl.Execute(&buf, map[string]interface{}{})
				if err != nil {
					out = "Execute: " + err.Error()
				} else {
					out = buf.String() // {"name1": "lmx", "no": 0.01 }
				}
			}
			if out != tt.out {
				t.Errorf("have: %s but want: %s", out, tt.out)
			} else {
				t.Log(out)
			}
		})
	}
}
