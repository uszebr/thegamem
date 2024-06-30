package chartcache

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChartCache_SetCache(t *testing.T) {
	cache := NewChartCache()

	gameID := "game1"
	boardQuantity := 10
	value := "chart1"

	cache.SetCache(gameID, boardQuantity, value)

	cachedValue, found := cache.GetCache(gameID, boardQuantity)
	assert.True(t, found)
	assert.Equal(t, value, cachedValue)
}

func TestChartCache_GetCache_NotFound(t *testing.T) {
	cache := NewChartCache()

	gameID := "game2"
	boardQuantity := 20

	value, found := cache.GetCache(gameID, boardQuantity)
	assert.False(t, found)
	assert.Empty(t, value)
}

func TestChartCache_GetCache_BoardQuantityMismatch(t *testing.T) {
	cache := NewChartCache()

	gameID := "game3"
	boardQuantity := 30
	value := "chart3"

	cache.SetCache(gameID, boardQuantity, value)

	mismatchedBoardQuantity := 40
	value, found := cache.GetCache(gameID, mismatchedBoardQuantity)
	assert.False(t, found)
	assert.Empty(t, value)
}

func TestChartCache_ConcurrentAccess(t *testing.T) {
	cache := NewChartCache()
	var wg sync.WaitGroup

	numRoutines := 100
	numIterations := 1000

	gameID := "game4"
	boardQuantity := 50
	value := "chart4"

	// Set cache in one goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.SetCache(gameID, boardQuantity, value)
	}()

	// Read cache in multiple goroutines
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				cachedValue, found := cache.GetCache(gameID, boardQuantity)
				if found {
					assert.Equal(t, value, cachedValue)
				}
			}
		}()
	}

	wg.Wait()
}
