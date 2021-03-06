package lru

import (
	"github.com/juntaki/transparent"
)

// NewCache returns LRUCache
func NewCache(bufferSize, cacheSize int) (transparent.Layer, error) {
	lru := NewStorage(cacheSize)
	layer, err := transparent.NewLayerCache(bufferSize, lru)
	if err != nil {
		return nil, err
	}
	return layer, nil
}
