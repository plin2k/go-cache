# go-cache

## Interface 

```go
    type Cache interface {
	Get(string) (interface{}, error)
	Set(string, interface{}, time.Duration) error
	Delete(string) error
	Flush() error
    }
```

## In Memory

```go
    mem, err := cache.NewMemory()
	if err != nil {
		fmt.Println(err)
	}
  
    mem.Set("test1", "hello", 0)
    mem.Get("test1")
    mem.Delete("test1")
    mem.Flush()
```

## Redis

```go
    redisDb0, err := cache.NewRedis("localhost:6379", "password", 0)
	if err != nil {
		fmt.Println(err)
	}
  
    redisDb0.Set("test1", "hello", 0)
    redisDb0.Get("test1")
    redisDb0.Delete("test1")
    redisDb0.Flush()
```

