package main

import "algorithm-go/offer"

func main() {
	lru := offer.Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)
}
