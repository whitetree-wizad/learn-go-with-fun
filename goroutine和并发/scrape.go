package main

import (
	"fmt"
	"sync"
	"time"
)

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) Visitlink(url string) int {
	//v.mu.Lock()
	//defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	test := &Visited{visited: map[string]int{}}
	go test.Visitlink("baidu.com")
	go test.Visitlink("baidu.com")
	go test.Visitlink("google.com")
	go test.Visitlink("libvio.fun")

	time.Sleep(2 * time.Second)
	test.mu.Lock()
	defer test.mu.Unlock()
	fmt.Println(test.visited)
}
