package cuckoofilter

import (
	"sync"
)

type CuckooFilter struct {
	filter *CFilter
	mutex  *sync.RWMutex
}

func New(opts ...option) *CuckooFilter {
	f := new(CuckooFilter)
	f.filter = NewCuckooFilter(opts...)
	f.mutex = new(sync.RWMutex)
	return f
}

func (f *CuckooFilter) Lookup(item []byte) bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.filter.Lookup(item)
}

func (f *CuckooFilter) Insert(item []byte) bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.filter.Insert(item)
}

func (f *CuckooFilter) Delete(item []byte) bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.filter.Delete(item)
}

func (f *CuckooFilter) Count() uint {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.filter.Count()
}

func (f *CuckooFilter) Update(filter *CFilter) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.filter = filter
	return
}
