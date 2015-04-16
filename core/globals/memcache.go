package globals

import (
	//corehelpers "collexy/core/helpers"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var Mc = SetupMemcacheServer()

func SetupMemcacheServer() (mc *memcache.Client) {
	mc = memcache.New("127.0.0.1:11211")
	log.Println("Memcache client started")
	return
}

//var Mc = corehelpers.StartMemcacheServer()
