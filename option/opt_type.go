package option

func (o *Opt) AddHandleNone(key interface{}, fn func() error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleNone(fn)
	})
}

func (o *Opt) HandleNone(key interface{}, fn func() error) (bool, error) {
	if o.Has(key) {
		return true, o.handleNone(fn)
	}

	return false, nil
}

func (o *Opt) handleNone(fn func() error) error {
	return fn()
}

func (o *Opt) AddHandleInt(key interface{}, fn func(val int) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleInt(key, fn)
	})
}

func (o *Opt) HandleInt(key interface{}, fn func(val int) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleInt(key, fn)
	}

	return false, nil
}

func (o *Opt) handleInt(key interface{}, fn func(val int) error) error {
	val, err := o.opts.GetInt(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleIntList(key interface{}, fn func(valList []int) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleIntList(key, fn)
	})
}

func (o *Opt) HandleIntList(key interface{}, fn func(valList []int) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleIntList(key, fn)
	}

	return false, nil
}

func (o *Opt) handleIntList(key interface{}, fn func(valList []int) error) error {
	val, err := o.opts.GetIntList(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleUint64(key interface{}, fn func(val uint64) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleUint64(key, fn)
	})
}

func (o *Opt) HandleUint64(key interface{}, fn func(val uint64) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleUint64(key, fn)
	}

	return false, nil
}

func (o *Opt) handleUint64(key interface{}, fn func(val uint64) error) error {
	val, err := o.opts.GetUint64(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleUint64List(key interface{}, fn func(valList []uint64) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleUint64List(key, fn)
	})
}

func (o *Opt) HandleUint64List(key interface{}, fn func(valList []uint64) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleUint64List(key, fn)
	}

	return false, nil
}

func (o *Opt) handleUint64List(key interface{}, fn func(valList []uint64) error) error {
	val, err := o.opts.GetUint64List(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleUint32(key interface{}, fn func(val uint32) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleUint32(key, fn)
	})
}

func (o *Opt) HandleUint32(key interface{}, fn func(val uint32) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleUint32(key, fn)
	}

	return false, nil
}

func (o *Opt) handleUint32(key interface{}, fn func(val uint32) error) error {
	val, err := o.opts.GetUint32(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleUint32List(key interface{}, fn func(valList []uint32) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleUint32List(key, fn)
	})
}

func (o *Opt) HandleUint32List(key interface{}, fn func(valList []uint32) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleUint32List(key, fn)
	}

	return false, nil
}

func (o *Opt) handleUint32List(key interface{}, fn func(valList []uint32) error) error {
	val, err := o.opts.GetUint32List(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleString(key interface{}, fn func(val string) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleString(key, fn)
	})
}

func (o *Opt) HandleString(key interface{}, fn func(val string) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleString(key, fn)
	}

	return false, nil
}

func (o *Opt) handleString(key interface{}, fn func(val string) error) error {
	val, err := o.opts.GetString(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleStringList(key interface{}, fn func(valList []string) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleStringList(key, fn)
	})
}

func (o *Opt) HandleStringList(key interface{}, fn func(valList []string) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleStringList(key, fn)
	}

	return false, nil
}

func (o *Opt) handleStringList(key interface{}, fn func(valList []string) error) error {
	val, err := o.opts.GetStringList(key)
	if err != nil {
		return err
	}

	return fn(val)
}

func (o *Opt) AddHandleJson(key interface{}, obj interface{}, fn func(val interface{}) error) *Opt {
	return o.AddHandle(key, func(key interface{}) error {
		return o.handleJson(key, obj, fn)
	})
}

func (o *Opt) HandleJson(key interface{}, obj interface{}, fn func(val interface{}) error) (bool, error) {
	if o.Has(key) {
		return true, o.handleJson(key, obj, fn)
	}

	return false, nil
}

func (o *Opt) handleJson(key interface{}, obj interface{}, fn func(val interface{}) error) error {
	err := o.opts.GetJson(key, &obj)
	if err != nil {
		return err
	}

	return fn(obj)
}
