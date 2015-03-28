package fixedhistory

import (
	"fmt"
	"sync"
)

type FixedArray struct {
	Mutex    *sync.Mutex
	items    []interface{}
	ValueMap func(interface{}) interface{}
}

func (f *FixedArray) Push(val interface{}) {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	f.items = append(f.items[1:], val)
}

func (f *FixedArray) Remove(val interface{}) error {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	index := -1
	for i, e := range f.items {
		if e == val {
			index = i
			break
		}
	}
	if index != -1 {
		f.items = append((f.items)[:index], (f.items)[index+1:]...)
	} else {
		return fmt.Errorf("val (%v) not found in (%v...)", val, f.items[0:5])
	}
	return nil
}

func (f *FixedArray) Contains(val interface{}) bool {
	for _, x := range f.items {
		if f.ValueMap != nil {
			x = f.ValueMap(x)
		}
		if x == val {
			return true
		}
	}
	return false
}

func (f *FixedArray) Get(val interface{}) interface{} {
	for _, x := range f.items {
		v := x
		if f.ValueMap != nil {
			v = f.ValueMap(x)
		}
		if v == val {
			return x
		}
	}
	return nil
}

func NewHistory(capacity int) *FixedArray {
	return &FixedArray{Mutex: &sync.Mutex{}, items: make([]interface{}, capacity)}
}

type CleanFn func(interface{}) bool

func (f *FixedArray) Cleanup(fn CleanFn) error {
	for _, item := range f.items {
		if fn(item) {
			if err := f.Remove(item); err != nil {
				fmt.Println("error cleaning up:", err) // debug
			}
		}
	}
	return nil
}
