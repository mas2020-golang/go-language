// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
  f     Func
  mu    sync.Mutex // mutex on cache
  cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
  value interface{}
  err   error
}

func New(f Func) *Memo {
  return &Memo{f: f, cache: make(map[string]result)}
}

// Get func is not concurrency-safe! You can access to the same memo.cache[key] from several go routines.
// you can test with:
// go clean -testcache && go test -v -run=TestConcurrent -race concurrency-example/memo1
func (memo *Memo) Get(key string) (interface{}, error) {
  memo.mu.Lock() // acquire the lock
  res, ok := memo.cache[key]
  if !ok {
    res.value, res.err = memo.f(key)
    memo.cache[key] = res
  }
  memo.mu.Unlock() // release the lock
  return res.value, res.err
}
