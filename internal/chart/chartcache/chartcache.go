package chartcache

import (
	"sync"
)

type ChartCache struct {
	cache map[string]cacheResult
	mutex sync.RWMutex
}

type cacheResult struct {
	boardQuantity int
	value         string
}

func NewChartCache() *ChartCache {
	return &ChartCache{
		cache: make(map[string]cacheResult),
	}
}

func (chartCache *ChartCache) SetCache(gameID string, boardQuantity int, value string) {
	chartCache.mutex.Lock()
	defer chartCache.mutex.Unlock()
	chartCache.cache[gameID] = cacheResult{
		boardQuantity: boardQuantity,
		value:         value,
	}
}

func (chartCache *ChartCache) GetCache(gameID string, boardQuantity int) (string, bool) {
	chartCache.mutex.RLock()
	defer chartCache.mutex.RUnlock()
	cached, found := chartCache.cache[gameID]
	if !found {
		return "", false
	}
	if cached.boardQuantity != boardQuantity {
		return "", false
	}
	return cached.value, true
}
