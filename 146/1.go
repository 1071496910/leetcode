package main

//Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
//
//Follow up:
//Could you do both operations in O(1) time complexity?
//
//Example:
//
//LRUCache cache = new LRUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

type LRUCache struct {
	l        *list.List
	h        map[int]*list.Element
	m        sync.Mutex
	capacity int
}

type kvPair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:        list.New(),
		h:        make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	this.m.Lock()
	defer this.m.Unlock()

	elem := this.h[key]
	if elem == nil {
		return -1
	}

	this.l.MoveToFront(elem)
	return elem.Value.(*kvPair).value

}

func (this *LRUCache) Put(key int, value int) {
	this.m.Lock()
	defer this.m.Unlock()

	kv := &kvPair{
		value: value,
		key:   key,
	}

	elem := this.h[key]
	if elem == nil {
		if this.l.Len() >= this.capacity {
			tail := this.l.Back()
			this.l.Remove(tail)
			delete(this.h, tail.Value.(*kvPair).key)
		}
	} else {
		this.l.Remove(elem)
	}

	elem = this.l.PushFront(kv)
	this.h[key] = elem
}
