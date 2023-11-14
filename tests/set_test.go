package tests

import (
	"github.com/onuragtas/go-cache/cache"
	"github.com/onuragtas/go-cache/cache/serializer"
	"log"
	"testing"
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
