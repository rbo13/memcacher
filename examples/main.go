package main

import (
	"log"

	"github.com/whaangbuu/memcacher"
)

func main() {
	m := memcacher.NewMemcached("localhost", "11211", "localhost:11211")

	ok, err := m.Set("test1", "chardy")

	if err != nil && !ok {
		log.Printf("ERROR DUE TO: %v", err)
	}

	log.Print(ok)
}
