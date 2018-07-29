package main

import (
	"fmt"
	"log"

	"github.com/whaangbuu/memcacher"
)

func main() {
	m := memcacher.NewMemcached("localhost", "11211", "localhost:11211")

	testKey := "test1"
	ok, err := m.Set(testKey, "chardy")

	if err != nil && !ok {
		log.Printf("ERROR DUE TO: %v", err)
	}

	log.Print(ok)
	log.Print("Memcache Value")

	val, _ := memcacher.GetVal(testKey, m)

	fmt.Print(val)
}
