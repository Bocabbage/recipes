package main

import (
	"sync"
	"time"
)

type Value struct {
	Val string
}

// todo: 限时🔒？

type ConcurrMap struct {
	m    map[string]Value
	lk   *sync.RWMutex
	cond *sync.Cond
	// timeout time.Duration
}

func (cm *ConcurrMap) Get(k string) string {
	cm.lk.RLock()
	defer cm.lk.RUnlock()

	for cm.m[k].Val == "" {
		cm.cond.Wait()
	}
	return cm.m[k].Val

}

func (cm *ConcurrMap) Put(k string, v string) bool {
	cm.lk.Lock()
	defer cm.lk.Unlock()

	cm.m[k] = Value{v}
	cm.cond.Broadcast()
	return true
}

func (cm *ConcurrMap) Drop(k string) bool {
	cm.lk.Lock()
	defer cm.lk.Unlock()

	delete(cm.m, k)
	return true
}

func NewConcurrMap(timeout time.Duration) ConcurrMap {
	lk := &sync.RWMutex{}
	return ConcurrMap{
		m:    make(map[string]Value),
		lk:   lk,
		cond: sync.NewCond(lk),
	}
}
