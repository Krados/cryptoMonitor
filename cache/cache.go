package cache

import "github.com/coocood/freecache"

var _default *freecache.Cache

func Init() {
	cacheSize := 300 * 1024 * 1024
	_default = freecache.NewCache(cacheSize)
}

func Get() *freecache.Cache {
	return _default
}
