package go_cache

import (
	"fmt"
	"time"
)

type Cache interface {
	Get(string) (interface{}, error)
	Set(string, interface{}, time.Duration) error
	Delete(string) error
	Flush() error
}

var errNotFound = fmt.Errorf("value not found")
