package lc146

import "container/list"

type LRU struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	represent map[int]*list.Element
	list      *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		represent: make(map[int]*list.Element),
		list:      list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.represent[key]
	if !ok {
		return -1
	}

	this.list.MoveToBack(e)
	return e.Value.(*LRU).value
}

func (this *LRUCache) Put(key int, value int)  {
	existElement, ok := this.represent[key]
	if ok {
		this.list.MoveToBack(existElement)
		existElement.Value.(*LRU).value = value
		return
	}

	if this.list.Len() >= this.capacity {
		oldestElement := this.list.Front()
		this.list.Remove(oldestElement)
		delete(this.represent, oldestElement.Value.(*LRU).key)
	}

	lru := &LRU{key: key, value: value}
	newElement := this.list.PushBack(lru)
	this.represent[key] = newElement
}
