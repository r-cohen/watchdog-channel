# watchdog-channel
A simple watchdog based on golang channels

## Usage
```go
// kicks at interval on the GetKickChannel()
w := NewWatchdog(time.Second * 3)
i := 0
for {
  select {
  case <-w.GetKickChannel():
    fmt.Println("kick!")
  }
}

// reset the watchdog timer
w.Reset()

// stop the watchdog permanently
w.Stop()
```

## Online example
[Go Playground](https://play.golang.org/p/wLtW6V4k8k7)
