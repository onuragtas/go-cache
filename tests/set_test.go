package tests

import (
	"log"
	"testing"

	"github.com/onuragtas/go-cache/cache"
	"github.com/onuragtas/go-cache/cache/serializer"
)

var cacheAdapter = cache.NewCacheAdapter([]string{"127.0.0.1:6379"})

func TestSet(t *testing.T) {

	var value = map[string]interface{}{
		"key": "value",
	}

	cErr := cacheAdapter.Set(serializer.NewSerializer(&serializer.Options{Serializer: serializer.NewJsonSerializer()}), "key", value, 0)
	if cErr != nil {
		log.Print(cErr)
	}

}

func TestDeleteHashWithPattern(t *testing.T) {
	cErr := cacheAdapter.DeleteHashWithPattern("*", "", 0, 1000)
	if cErr != nil {
		log.Print(cErr)
	}

}
