# Concurrency example

In this example we can test the concurrency and analyze the options to avoid it.

The project is a go module and the main parts are:

- `memotest`: it is the executor of the call in a sequential and concurrent way
- `memo`: contains the type that memorizes the results of the call. Memo can store
the result of a call in a specific map. You can pass to the Memo constructor your own
  func, memo will call it and store the result internally. If you get again the same
  func with the same key, memo will immediately give you the cached result.
  
## Versions

### memo1
It is the first version, and it is not concurrency safe. You can call it for testing as:
  ```shell
  cd goroutines/race-conditions/concurrency-example
  go clean -testcache && go test -v -run=TestConcurrent -race concurrency-example/memo1
  ```
  You notice that there are sequential and concurrent calls, in the test above you call only the concurrent
  part and Go will notice you that you have a data race. The problem is here:
  ```go
  memo.cache[key] = res
  ```
  in the Get method of the Memo type. More goroutines can access and modify this map not safely.

### memo2
It is the second version, and it is concurrency safe. We lose the performance gain because to avoid the data race
we added a lock on the entire Get func:
```go
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
```

you can test as:

```shell
  cd goroutines/race-conditions/concurrency-example
  go clean -testcache && go test -v -run=TestConcurrent -race concurrency-example/memo2
  ```
***Problem***: Get serializes all the I/O operations we intended to parallelize. The advantage of the go routines
is lost.

### memo3
It is the third version, and it is concurrency safe. We lose the performance gain because to avoid the data race
we added a lock on the entire Get func:
```go
func (memo *Memo) Get(key string) (interface{}, error) {
  memo.mu.Lock() // acquire the lock
  res, ok := memo.cache[key]
  memo.mu.Unlock() // release the lock
  if !ok {
    res.value, res.err = memo.f(key)
    memo.mu.Lock() // acquire the lock
    memo.cache[key] = res
    memo.mu.Unlock() // release the lock
  }
  return res.value, res.err
}
```
The performance it's fine but we risk to do the same thing more times. If two different goroutines with the same key happen
at the same time the get an empty value from the cache and invoke both the HTTP get. Then one will override the data written
by the other.

### memo4

This example is more complex but avoid the problem of the [above](#memo3).
Now the `cache` variable of the Memo struct is:
```go
cache map[string]*entry
```

The entry contains the result, and a broadcast channel to inform the other go routines that this entry is available:
```go
type entry struct {
  res   result
  ready chan struct{} // closed when res is ready
}
```
The `Get` func has the most logic:
```go
func (memo *Memo) Get(key string) (value interface{}, err error) {
  memo.mu.Lock()
  e := memo.cache[key]
  if e == nil {
    // This goroutine becomes responsible for computing
    // the value and broadcasting the ready condition.
    e = &entry{ready: make(chan struct{})}
    memo.cache[key] = e
    memo.mu.Unlock()
  
    e.res.value, e.res.err = memo.f(key)
    
    close(e.ready) // broadcast ready condition
  } else {
    // This is a repeat request for this key.
    memo.mu.Unlock()
    
    <-e.ready // wait for ready condition
  }
  return e.res.value, e.res.err
}
```
The lock is released:

- after the creation of the entry in the cache table
- if the cache already exist

By this way the HTTP request is available after that the cache has been created. The other routines on the same key will find the
key already present in the cache and will wait until the channel is closed. No more HTTP calls on the same key will be
executed.
To test the `memo4` solution type:

```shell
go clean -testcache && go test -v -run=TestConcurrent -race concurrency-example/memo4
```