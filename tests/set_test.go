package tests

import (
	"log"
	"testing"

	"github.com/onuragtas/go-cache/cache"
	"github.com/onuragtas/go-cache/cache/serializer"
)

func TestSet(t *testing.T) {
	cacheAdapter := cache.NewCacheAdapter([]string{"192.168.36.240:6379"})

	var value = map[string]interface{}{
		"key": "value",
	}

	cErr := cacheAdapter.Set(serializer.NewSerializer(&serializer.Options{Serializer: serializer.NewJsonSerializer()}), "key", value, 0)
	if cErr != nil {
		log.Print(cErr)
	}

}

func TestDeleteHashWithPattern(t *testing.T) {
	cacheAdapter := cache.NewCacheAdapter([]string{"192.168.36.240:6379"})

	cErr := cacheAdapter.DeleteHashWithPattern("comment:{upid}:nc_96638713_*", "*", 0, 1000)
	if cErr != nil {
		log.Print(cErr)
	}

}
