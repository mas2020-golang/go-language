# Channels

## Close a channel

In case the sender knows that there are no more data to send elaborate it can close the channel.

```golang
close(chan)
```

Any further operation will cause a ***panic***. After the closed channel has been drained the subsequent receive
operation will go without blocking but will yeld a value of zero in an infinite loop. To **check** if a channel has been
closed use this syntax:

```golang
for {
  x, ok := <-channel
  if !ok {
    break // channel was closed and drained
  }
  ...
}
```
In the example above if a channel is empty the program will stop reading from it. An alternative is to use
the `range` loop:
```go
for {
    x, ok := <-naturals
    if !ok {
        break // channel was closed and drained
    }
    squares <- x * x
}
```


