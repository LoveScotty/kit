package option

import (
	"sort"

	"scotty/kit/util"
)

type Option struct {
	Key int32
	Val string
}

type OptFn func(key interface{}) error

type (
	optHandle struct {
		key  int32
		fn   OptFn
		sort int
	}
	optHandleList []optHandle
)

type optHandleSortBySort struct {
	optHandleList
}

func (o optHandleList) Len() int      { return len(o) }
func (o optHandleList) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o optHandleSortBySort) Less(i, j int) bool {
	return o.optHandleList[i].sort < o.optHandleList[j].sort
}

type Opt struct {
	opts    *Opts
	handles map[int32]optHandle
	val     int
}

func NewOpt(opts ...*Option) *Opt {
	var o Opt

	o.opts = NewOpts(opts...)
	o.handles = make(map[int32]optHandle, o.opts.Len())

	return &o
}

func (o *Opt) Exec() error {
	for _, v := range o.getSortedHandles() {
		if !o.Has(v.key) {
			continue
		}
		err := v.fn(v.key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Opt) Opts() *Opts {
	return o.opts
}

func (o *Opt) Has(key interface{}) bool {
	return o.Opts().Has(key)
}

func (o *Opt) HasAllOf(keys ...interface{}) bool {
	for _, v := range keys {
		if !o.Has(v) {
			return false
		}
	}

	return true
}

func (o *Opt) HasOneOf(keys ...interface{}) bool {
	for _, v := range keys {
		if o.Has(v) {
			return true
		}
	}

	return false
}

func (o *Opt) getSortedHandles() optHandleList {
	var list optHandleList
	for _, v := range o.handles {
		list = append(list, v)
	}
	sort.Sort(optHandleSortBySort{list})

	return list
}

func (o *Opt) getSortVal() int {
	o.val += 1

	return o.val
}

func (o *Opt) AddHandle(key interface{}, fn OptFn) *Opt {
	k := util.ToInt32(key)
	o.handles[k] = optHandle{
		key:  k,
		fn:   fn,
		sort: o.getSortVal(),
	}

	return o
}
