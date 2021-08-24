package memtable

import "time"

func NewWatcher(cache *LruCache) error {
	Oldest := cache.tail
	if Oldest.TimeOut == 0 || Oldest.start.Add(time.Duration(Oldest.TimeOut)).After(time.Now()) {
		cache.tail = cache.tail.pre
	}

	return nil
}
