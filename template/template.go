package template

import (
	"bytes"
	"fmt"
	htmltmpl "html/template"
	"reflect"
	"strings"
	texttmpl "text/template"
	"text/template/parse"
)

type Template interface {
	DefinedTemplates() string
	Name() string
}

// Parse 模板解析
func Parse(t Template, text string) (err error) {
	err = bindFunc(t, t.Name(), text)
	if err != nil {
		return
	}

	switch tp := t.(type) {
	case *texttmpl.Template:
		_, err = tp.Parse(text)
	case *htmltmpl.Template:
		_, err = tp.Parse(text)
	}
	if err != nil {
		return
	}

	return
}

// bindFunc 绑定方法到template中
func bindFunc(t Template, name, text string) error {
	var leftDelim, rightDelim string
	switch tp := t.(type) {
	case *texttmpl.Template:
		leftDelim = reflect.ValueOf(tp).Elem().FieldByName("leftDelim").String()
		rightDelim = reflect.ValueOf(tp).Elem().FieldByName("rightDelim").String()
	case *htmltmpl.Template:
		leftDelim = reflect.ValueOf(tp).Elem().FieldByName("text").Elem().FieldByName("leftDelim").String()
		rightDelim = reflect.ValueOf(tp).Elem().FieldByName("text").Elem().FieldByName("rightDelim").String()
	default:
		return fmt.Errorf("template type: %T is not support yet", tp)
	}

	treeSet := make(map[string]*parse.Tree)
	tree := parse.New(name)
	tree.Mode = parse.SkipFuncCheck
	_, err := tree.Parse(text, leftDelim, rightDelim, treeSet)
	if err != nil {
		return err
	}

	funcMap := make(map[string]interface{})
	for name := range treeSet {
		if err := addFunc(t, name, funcMap); err != nil {
			return err
		}
	}

	switch tp := t.(type) {
	case *texttmpl.Template:
		tp.Funcs(funcMap)
	case *htmltmpl.Template:
		tp.Funcs(funcMap)
	}

	return nil
}

func addFunc(t Template, name string, funcMap map[string]interface{}) error {
	fn, bundle, err := bundler(name)
	if err != nil {
		return err
	}
	if len(fn) == 0 {
		return nil
	}
	switch tp := t.(type) {
	case *texttmpl.Template:
		funcMap[fn] = func(args ...interface{}) (string, error) {
			t := tp.Lookup(name)
			if t == nil {
				return "", fmt.Errorf("lost template %q", name)
			}
			arg, err := bundle(args)
			if err != nil {
				return "", err
			}
			var buf bytes.Buffer
			err = t.Execute(&buf, arg)
			if err != nil {
				return "", err
			}
			return buf.String(), nil
		}
	case *htmltmpl.Template:
		funcMap[fn] = func(args ...interface{}) (htmltmpl.HTML, error) {
			t := tp.Lookup(name)
			if t == nil {
				return "", fmt.Errorf("lost template %q", name)
			}
			arg, err := bundle(args)
			if err != nil {
				return "", err
			}
			var buf bytes.Buffer
			err = t.Execute(&buf, arg)
			if err != nil {
				return "", err
			}
			return htmltmpl.HTML(buf.String()), nil
		}
	}

	return nil
}

func bundler(name string) (fn string, bundle func(args []interface{}) (interface{}, error), err error) {
	f := strings.Fields(name)
	if len(f) == 0 {
		return
	}

	fn = f[0]
	if len(f) == 1 {
		bundle = func(args []interface{}) (interface{}, error) {
			if len(args) == 0 {
				return nil, nil
			}
			if len(args) == 1 {
				return args[0], nil
			}
			return nil, fmt.Errorf("too many arguments in call to template %s", fn)
		}
		return
	}

	var sawQ bool
	for i, argName := range f[1:] {
		if strings.HasSuffix(argName, "...") {
			if i != len(f)-2 {
				err = fmt.Errorf("invalid template name %q: %s is not last argument", name, argName)
				return
			}
			break
		}
		if strings.HasSuffix(argName, "?") {
			sawQ = true
			continue
		}
		if sawQ {
			err = fmt.Errorf("invalid template name %q: required %s after optional %s", name, argName, f[i])
			return
		}
	}

	bundle = func(args []interface{}) (interface{}, error) {
		m := make(map[string]interface{})
		for _, argName := range f[1:] {
			if strings.HasSuffix(argName, "...") {
				m[strings.TrimSuffix(argName, "...")] = args
				args = nil
				break
			}
			if strings.HasSuffix(argName, "?") {
				prefix := strings.TrimSuffix(argName, "?")
				if len(args) == 0 {
					m[prefix] = nil
				} else {
					m[prefix], args = args[0], args[1:]
				}
				continue
			}
			if len(args) == 0 {
				return nil, fmt.Errorf("too few arguments in call to template %s", fn)
			}
			m[argName], args = args[0], args[1:]
		}
		if len(args) > 0 {
			return nil, fmt.Errorf("too many arguments in call to template %s", fn)
		}
		return m, nil
	}

	return
}
