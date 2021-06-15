package leetcode

import "container/list"

type Cache struct {
	// 内存总空间
	maxBytes int64
	// 已存储空间
	nbytes int64
	// 双向链表
	ll    *list.List
	cache map[string]*list.Element
	// 在清除内存时执行
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

var _ Value = &Cache{}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Add(key string, value Value) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToBack(element)
	} else {
		// 增加内存
		c.nbytes += int64(value.Len()) + int64(len(key))
		// 队尾增加
		c.ll.PushBack(&entry{key, value})
		c.cache[key] = c.ll.Back()
	}

	// if c.nbytes > c.maxBytes {
	// 	c.RemoveOldest()
	// }
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToBack(element)
		res := element.Value.(*entry)
		return res.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	element := c.ll.Front()
	if element == nil {
		return
	}
	c.ll.Remove(element)
	front := element.Value.(*entry)
	c.nbytes -= int64(front.value.Len()) + int64(len(front.key))
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
