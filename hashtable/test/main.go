package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/lehaisonmath6/Algorithm/hashtable"
)

func main() {
	h := hashtable.NewHashTable(10)
	rand.Seed(time.Now().Unix())
	var lsKeys []int64
	for i := 0; i < 10; i++ {
		k := rand.Int63()
		v := rand.Int63()
		ok := h.Put(k, v)
		if !ok {
			log.Println("[ERROR] put k", k, "v", v, "err")
			continue
		}
		log.Println("put k", k, "v", v)
		lsKeys = append(lsKeys, k)

	}
	log.Println("list key", lsKeys, "len", len(lsKeys))
	for _, k := range lsKeys {
		v, ok := h.Get(k)
		if !ok {
			log.Println("[ERRPR] get k", k, "not found")
			continue
		}
		log.Println("k", k, "v", v)
	}
}
