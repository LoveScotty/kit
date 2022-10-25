package option

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"scotty/kit/util"
)

func NewNoneOpt(key interface{}) *Option {
	return &Option{
		Key: util.ToInt32(key),
	}
}

func NoneOpt(key interface{}) OptFunc {
	return func() *Option {
		return NewNoneOpt(key)
	}
}

func NewIntOpt(key interface{}, val int) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: strconv.Itoa(val),
	}
}

func IntOpt(key interface{}, val int) OptFunc {
	return func() *Option {
		return NewIntOpt(key, val)
	}
}

func (o *Opts) GetInt(key interface{}) (int, error) {
	str, err := o.Val(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(str)
}

func NewIntListOpt(key interface{}, val ...int) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(val)), ","), "[]"),
	}
}

func IntListOpt(key interface{}, val ...int) OptFunc {
	return func() *Option {
		return NewIntListOpt(key, val...)
	}
}

func (o *Opts) GetIntList(key interface{}) ([]int, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return nil, err
	}

	return util.Str2IntList(valStr, ",")
}

func NewUint64Opt(key interface{}, val uint64) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: fmt.Sprintf("%d", val),
	}
}

func Uint64Opt(key interface{}, val uint64) OptFunc {
	return func() *Option {
		return NewUint64Opt(key, val)
	}
}

func (o *Opts) GetUint64(key interface{}) (uint64, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(valStr, 10, 64)
}

func NewUint64ListOpt(key interface{}, val ...uint64) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(val)), ","), "[]"),
	}
}

func Uint64ListOpt(key interface{}, val ...uint64) OptFunc {
	return func() *Option {
		return NewUint64ListOpt(key, val...)
	}
}

func (o *Opts) GetUint64List(key interface{}) ([]uint64, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return nil, err
	}

	return util.Str2U64List(valStr, ",")
}

func NewUint32Opt(key interface{}, val uint32) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: fmt.Sprintf("%d", val),
	}
}

func Uint32Opt(key interface{}, val uint32) OptFunc {
	return func() *Option {
		return NewUint32Opt(key, val)
	}
}

func (o *Opts) GetUint32(key interface{}) (uint32, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return 0, err
	}
	u64, err := strconv.ParseUint(valStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(u64), nil
}

func NewUint32ListOpt(key interface{}, val ...uint32) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(val)), ","), "[]"),
	}
}

func Uint32ListOpt(key interface{}, val ...uint32) OptFunc {
	return func() *Option {
		return NewUint32ListOpt(key, val...)
	}
}

func (o *Opts) GetUint32List(key interface{}) ([]uint32, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return nil, err
	}

	return util.Str2U32List(valStr, ",")
}

func NewStringOpt(key interface{}, val string) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: val,
	}
}

func StringOpt(key interface{}, val string) OptFunc {
	return func() *Option {
		return NewStringOpt(key, val)
	}
}

func (o *Opts) GetString(key interface{}) (string, error) {
	return o.Val(key)
}

func NewStringListOpt(key interface{}, val ...string) *Option {
	return &Option{
		Key: util.ToInt32(key),
		Val: strings.Join(val, ","),
	}
}

func StringListOpt(key interface{}, val ...string) OptFunc {
	return func() *Option {
		return NewStringListOpt(key, val...)
	}
}

func (o *Opts) GetStringList(key interface{}) ([]string, error) {
	valStr, err := o.Val(key)
	if err != nil {
		return nil, err
	}

	return strings.Split(valStr, ","), nil
}

func NewJsonOpt(key, val interface{}) *Option {
	if val == nil {
		return &Option{
			Key: util.ToInt32(key),
		}
	}
	js, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return &Option{
		Key: util.ToInt32(key),
		Val: string(js),
	}
}

func (o *Opts) GetJson(key, val interface{}) error {
	value, err := o.Val(key)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(value), val)
}

func JsonOpt(key, val interface{}) OptFunc {
	return func() *Option {
		return NewJsonOpt(key, val)
	}
}
