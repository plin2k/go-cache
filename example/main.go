package main

import (
	"fmt"
	cache "github.com/plin2k/go-cache"
	"time"
)

func main() {
	memory()
	redis()
}

func redis() {
	red, err := cache.NewRedis("127.0.0.1:6379", "", 0)
	if err != nil {
		fmt.Println(err)
	}

	test(red)
}

func memory() {
	mem, err := cache.NewMemory()
	if err != nil {
		fmt.Println(err)
	}

	test(mem)
}

func test(myCache cache.Cache) {
	fmt.Println("Set test1 - ", myCache.Set("test1", "hello", 0))
	fmt.Println("Set test2 - ", myCache.Set("test2", struct {
		Name string
		Age  int
	}{"Alex", 10}, 2*time.Second))

	fmt.Print("Get test1 - ")
	fmt.Println(myCache.Get("test1"))

	time.Sleep(time.Second * 2)

	fmt.Print("Get test2 - ")
	fmt.Println(myCache.Get("test2"))

	fmt.Println("Delete - ", myCache.Delete("test2"))
	fmt.Println("Flush - ", myCache.Flush())

}
