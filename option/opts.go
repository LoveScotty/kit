package option

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"scotty/kit/util"
)

var ErrOptKeyNotFound = errors.New("key not found")

type OptFunc func() *Option

type Opts struct {
	opts map[int32]*Option
}

func NewOpts(opts ...*Option) *Opts {
	var o Opts
	o.opts = make(map[int32]*Option)
	o.Set(opts...)
	return &o
}

func (o *Opts) Set(opts ...*Option) *Opts {
	for _, v := range opts {
		if v == nil {
			continue
		}
		o.opts[v.Key] = v
	}
	return o
}

func (o *Opts) SetIf(needSet bool, opts ...*Option) *Opts {
	if !needSet {
		return o
	}
	for _, v := range opts {
		if v == nil {
			continue
		}
		o.opts[v.Key] = v
	}
	return o
}

func (o *Opts) Has(key interface{}) bool {
	_, ok := o.opts[util.ToInt32(key)]
	return ok
}

func (o *Opts) Len() int {
	return len(o.opts)
}

func (o *Opts) Val(key interface{}) (string, error) {
	if o.Has(key) {
		return o.opts[util.ToInt32(key)].Val, nil
	}
	return "", ErrOptKeyNotFound
}

func (o *Opts) List() []*Option {
	var list []*Option
	for _, v := range o.opts {
		list = append(list, v)
	}
	return list
}

func (o *Opts) KeyString() string {
	var keyList []int
	for _, v := range o.opts {
		keyList = append(keyList, int(v.Key))
	}
	sort.Ints(keyList)
	var strList []string
	for _, v := range keyList {
		strList = append(strList, strconv.Itoa(v))
	}
	return strings.Join(strList, ",")
}

func GenOpts(constructors ...OptFunc) *Opts {
	o := NewOpts()
	for _, v := range constructors {
		o.Set(v())
	}
	return o
}

func GenOptsList(constructors ...OptFunc) []*Option {
	return GenOpts(constructors...).List()
}

func GenNoneOptsList(keys ...interface{}) []*Option {
	var list []*Option
	for _, v := range keys {
		list = append(list, &Option{
			Key: util.ToInt32(v),
		})
	}
	return list
}
