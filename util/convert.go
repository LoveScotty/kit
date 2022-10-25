package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToInt(key interface{}) int {
	switch reflect.TypeOf(key).Kind() {
	case reflect.Int,
		reflect.Int32,
		reflect.Int64,
		reflect.Int16,
		reflect.Int8:
		return int(reflect.ValueOf(key).Int())
	case reflect.Uint,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uint16,
		reflect.Uint8:
		return int(reflect.ValueOf(key).Uint())
	default:
		panic("类型错误")
	}
	return 0
}

func ToInt32(key interface{}) int32 {
	switch reflect.TypeOf(key).Kind() {
	case reflect.Int,
		reflect.Int32,
		reflect.Int64,
		reflect.Int16,
		reflect.Int8:
		return int32(reflect.ValueOf(key).Int())
	case reflect.Uint,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uint16,
		reflect.Uint8:
		return int32(reflect.ValueOf(key).Uint())
	default:
		panic(fmt.Sprintf("类型错误: %T", key))
	}
	return 0
}

func Str2IntList(s, sep string) ([]int, error) {
	var err error
	list := strings.Split(s, sep)
	list2 := make([]int, len(list))
	for i, v := range list {
		list2[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return list2, nil
}

func Str2U32List(s, sep string) ([]uint32, error) {
	list := strings.Split(s, sep)
	u32List := make([]uint32, len(list))
	for i, v := range list {
		u64, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, err
		}
		u32List[i] = uint32(u64)
	}
	return u32List, nil
}

func Str2U64List(s, sep string) ([]uint64, error) {
	list := strings.Split(s, sep)
	u64List := make([]uint64, len(list))
	for i, v := range list {
		u64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		u64List[i] = u64
	}
	return u64List, nil
}
