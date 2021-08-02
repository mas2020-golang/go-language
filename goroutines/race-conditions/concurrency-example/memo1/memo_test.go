package memo_test

import (
  memotest "concurrency-example"
  "concurrency-example/memo1"
  "testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
  m := memo.New(httpGetBody)
  memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
 m := memo.New(httpGetBody)
 memotest.Concurrent(t, m)
}
