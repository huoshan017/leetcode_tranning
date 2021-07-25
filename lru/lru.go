package lru

type LRUCache struct {
	data      []int       // 数据列表
	head      int         // 头节点索引
	count     int         // 节点数
	key2Index map[int]int // key对应到索引
	index2Key map[int]int // 索引对应到key
}

func Constructor(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("invalid capcity")
	}
	return &LRUCache{
		data:      make([]int, capacity),
		key2Index: make(map[int]int),
		index2Key: make(map[int]int),
	}
}

func (l *LRUCache) Get(key int) int {
	// key对应的索引
	idx, o := l.key2Index[key]
	if !o {
		return -1
	}
	// 索引非法
	if idx >= len(l.data) {
		return -1
	}

	// 非头节点索引
	if idx != l.head {
		// 交换该值到最前面
		tmp := l.data[l.head]
		l.data[l.head] = l.data[idx]
		l.data[idx] = tmp
		// 更新两个值的key和index的关系
		l.key2Index[key] = l.head
		k := l.index2Key[l.head]
		l.index2Key[l.head] = key
		l.key2Index[k] = idx
		l.index2Key[idx] = k
	}
	return l.data[l.head]
}

func (l *LRUCache) Put(key int, value int) {
	// 已存在
	if idx, o := l.key2Index[key]; o {
		l.data[idx] = value
		return
	}

	// 没有数据，头节点在数组最后
	if l.isEmpty() {
		l.head = len(l.data) - 1
	} else {
		if l.head > 0 {
			l.head -= 1
		} else {
			l.head = len(l.data) - 1
		}
		// 上面指定了新节点的插入位置索引，如果队列已满，下面做插入位置的key删除
		if l.isFull() {
			if v, o := l.index2Key[l.head]; o {
				delete(l.key2Index, v)
			}
		}
	}

	// 新节点
	l.data[l.head] = value
	l.key2Index[key] = l.head
	l.index2Key[l.head] = key

	if !l.isFull() {
		l.count += 1
	}
}

func (l *LRUCache) isEmpty() bool {
	return l.count <= 0
}

func (l *LRUCache) isFull() bool {
	return l.count >= len(l.data)
}

type LRU struct {
}
