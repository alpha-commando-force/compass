package memtable

import (
	"sync"
	"time"

	"github.com/cornelk/hashmap"
)

type LruCache struct {
	mu         sync.Mutex
	size       int
	survive    int
	cache      *hashmap.HashMap
	head, tail *LruNode
}

type LruNode struct {
	NodeIP    string
	value     []byte
	TimeOut   int64
	start     time.Time
	pre, next *LruNode
}

func NewNode(nodeIP string, value []byte, timeOut int64) *LruNode {
	return &LruNode{
		NodeIP:  nodeIP,
		value:   value,
		TimeOut: timeOut,
		start:   time.Now(),
	}
}

func initCache() *LruCache {
	cache := &LruCache{
		cache: &hashmap.HashMap{},
		head:  NewNode("127.0.0.0:0", nil, 0),
		tail:  NewNode("127.0.0.0:1", nil, 0),
	}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head

	return cache
}

func (cache *LruCache) AddNode(node *LruNode) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.addToHead(node)
}

func (cache *LruCache) moveToHead(node *LruNode) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.DeleteNode(node)
	cache.addToHead(node)
}

func (cache *LruCache) addToHead(node *LruNode) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.head.next.pre = node
	node.next = cache.head.next
	cache.head.next = node
	node.pre = cache.head
}

func (cache *LruCache) DeleteNode(node *LruNode) {
	_, ok := cache.cache.Get(node.NodeIP)
	if ok {
		cache.cache.Del(node.NodeIP)
	}
}
