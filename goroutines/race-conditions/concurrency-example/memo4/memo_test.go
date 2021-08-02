package memo_test

import (
  "concurrency-example/memo4"
  "testing"
)

var httpGetBody = memo.HTTPGetBody

func Test(t *testing.T) {
  m := memo.New(httpGetBody)
  memo.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
 m := memo.New(httpGetBody)
 memo.Concurrent(t, m)
}
