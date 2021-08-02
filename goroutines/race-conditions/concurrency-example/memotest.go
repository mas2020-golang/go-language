package concurrency_example

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "sync"
  "testing"
  "time"
)

func httpGetBody(url string) (interface{}, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
  ch := make(chan string)
  go func() {
    for _, url := range []string{
      "https://golang.org",
      "https://godoc.org",
      "https://play.golang.org",
      "https://www.microsoft.com",
      "https://golang.org",
      "https://godoc.org",
      "https://play.golang.org",
      "https://www.microsoft.com",
    } {
      ch <- url
    }
    close(ch)
  }()
  return ch
}

type M interface {
  Get(key string) (interface{}, error)
}

// Sequential calls the m.Get func passing the url read from the channel returned by the
// incomingURLs func. It waits the end of all the go routines before returns to the caller.
func Sequential(t *testing.T, m M) {
  for url := range incomingURLs() {
    start := time.Now()
    value, err := m.Get(url)
    if err != nil {
      log.Print(err)
    }
    fmt.Printf("%s, %s, %d bytes\n",
      url, time.Since(start), len(value.([]byte)))
  }
}

// Concurrent calls (in async way using n goroutines) the m.Get func passing the url read from the channel returned by the
// incomingURLs func. It waits the end of all the go routines before returns to the caller.
func Concurrent(t *testing.T, m M) {
  var n sync.WaitGroup
  for url := range incomingURLs() {
    n.Add(1)
    go func(url string) {
      start := time.Now()
      value, err := m.Get(url)
      if err != nil {
        log.Print(err)
      }
      fmt.Printf("%s, %s, %d bytes\n",
        url, time.Since(start), len(value.([]byte)))
      n.Done()
    }(url)
  }
  n.Wait()
}
