package main

import "fmt"

type EvictionAlgorithm interface {
	evict(c *Cache)
}

type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	//TODO remove space
	fmt.Println("Evicting by fifo strategy")
}

type Lru struct{}

func (l Lru) evict(c *Cache) {
	//TODO remove space
	fmt.Println("Evicting by lru strategy")
}

type Lfu struct{}

func (l Lfu) evict(c *Cache) {
	//TODO remove space
	fmt.Println("Evicting by lfu strategy")
}

type Cache struct {
	storage           map[string]string
	evictionAlgorithm EvictionAlgorithm
	capacity          int
	maxCapacity       int
}

func (c *Cache) setEvicitionAlgorithm(algorithm EvictionAlgorithm) {
	c.evictionAlgorithm = algorithm
}

func initCache(e EvictionAlgorithm) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:           storage,
		evictionAlgorithm: e,
		capacity:          0,
		maxCapacity:       2,
	}
}

func (c *Cache) evict() {
	c.evictionAlgorithm.evict(c)
	c.capacity--
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvicitionAlgorithm(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvicitionAlgorithm(fifo)

	cache.add("e", "5")
	fmt.Println(cache.storage)
	fmt.Println(cache.capacity)
	fmt.Println(cache.maxCapacity)
}
