package option

import (
	"testing"
)

func TestOptExec(t *testing.T) {
	type TestJson struct {
		A int `json:"a"`
	}
	options := GenOptsList(
		NoneOpt(1),
		IntOpt(2, 2),
		IntListOpt(3, 3, 3, 3),
		Uint32Opt(4, 4),
		Uint32ListOpt(5, 5, 5, 5),
		Uint64Opt(6, 6),
		Uint64ListOpt(7, 7, 7, 7),
		StringOpt(8, "hello"),
		StringListOpt(9, "hello", "hi"),
		JsonOpt(10, &TestJson{
			A: 6,
		}),
	)
	optX := NewOpt(options...)
	err := optX.
		AddHandleNone(1, func() error {
			t.Logf("HandleNone: %v", 1)
			return nil
		}).
		AddHandleInt(2, func(val int) error {
			t.Logf("HandleInt: %v", val)
			return nil
		}).
		AddHandleIntList(3, func(val []int) error {
			t.Logf("HandleIntList: %v", val)
			return nil
		}).
		AddHandleUint32(4, func(val uint32) error {
			t.Logf("HandleUint32: %v", 4)
			return nil
		}).
		AddHandleUint32List(5, func(val []uint32) error {
			t.Logf("HandleUint32List: %v", val)
			return nil
		}).
		AddHandleUint64(6, func(val uint64) error {
			t.Logf("HandleUint64: %v", val)
			return nil
		}).
		AddHandleUint64List(7, func(val []uint64) error {
			t.Logf("HandleUint64List: %v", val)
			return nil
		}).
		AddHandleString(8, func(val string) error {
			t.Logf("HandleString: %v", val)
			return nil
		}).
		AddHandleStringList(9, func(val []string) error {
			t.Logf("HandleStringList: %v", val)
			return nil
		}).
		AddHandleJson(10, &TestJson{}, func(val interface{}) error {
			t.Logf("HandleJson: %v", val)
			return nil
		}).
		Exec()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
}
