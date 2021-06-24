/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package leetcode

import "container/list"

// @lc code=start

type LRUCache struct {
	// 总内存空间
	maxBytes int64
	// 已存储空间
	nbytes int64
	// 双向链表 通过双向链表建立队列实现 最新最少访问
	ll *list.List
	// 存储空间map结构 值是双向链表中对应节点的指针
	cache map[int]*list.Element
}

//  ll 中数据类型 在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   int
	value int
}

// //
// type Value interface {
// 	Len() int
// }

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxBytes: int64(capacity),
		nbytes:   0,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		this.ll.MoveToBack(value)
		entry := value.Value.(*entry)
		return entry.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		//  可能存在相同的key 对应不同的value
		en := v.Value.(*entry)
		en.value = value
		this.ll.MoveToBack(v)
	} else {
		ele := this.ll.PushBack(&entry{key, value})
		this.cache[key] = ele
		this.nbytes++
	}

	for this.maxBytes != 0 && this.maxBytes < this.nbytes {
		this.RemoveOldest()
	}
}

func (this *LRUCache) RemoveOldest() {
	ele := this.ll.Front()
	if this.ll.Len() > 0 && ele != nil {
		this.ll.Remove(ele)
		eleValue := ele.Value.(*entry)
		delete(this.cache, eleValue.key)
		this.nbytes--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
