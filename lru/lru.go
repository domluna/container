package lru

import (
	"container/list"
	"sync"
)

// Cache is a LRU Cache. The cache is concurrent friendly.
type Cache struct {
	list       *list.List
	maxEntries int
	cache      map[string]*list.Element

	lock sync.Mutex
}

type elem struct {
	key   string
	value interface{}
}

// New creates a LRU Cache
func New(maxEntries int) *Cache {
	return &Cache{
		list:       list.New(),
		maxEntries: maxEntries,
		cache:      make(map[string]*list.Element),
	}
}

// Add inserts element key,value pair into the cache.
//
// If the key is exists the element becomes the most recently
// used element and the previously value is overwritten.
//
// Otherwise the element is added and becomes the most recently used
// element.
func (c *Cache) Add(key string, v interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	e, ok := c.cache[key]
	// Already in, bump it up to the front
	if ok {
		c.list.MoveToFront(e)
		e.Value.(*elem).value = v
		return
	}
	ee := c.list.PushFront(&elem{key, v})
	c.cache[key] = ee

	if c.list.Len() > c.maxEntries {
		c.remove()
	}
}

// Get retrieves the element value given a key. If the cache does not contain
// the key, (nil, false) is returned.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.(*elem).value, true
	}
	return nil, false
}

// removeOldest deletes the least recently used element in the cache.
func (c *Cache) removeOldest() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.remove()
}

// Len returns the number of elements in the cache.
func (c *Cache) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.list.Len()
}

func (c *Cache) remove() {
	e := c.list.Back()
	if e == nil {
		return
	}
	c.list.Remove(e)
	delete(c.cache, e.Value.(*elem).key)
}
