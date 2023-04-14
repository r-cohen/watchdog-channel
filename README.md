# watchdog-channel
A simple watchdog timer based on golang channels

## Installation
`go get github.com/r-cohen/watchdog-channel`

## Usage
```go
// kicks at interval on the GetKickChannel()
wd := watchdog.NewWatchdog(time.Second * 3)
go func(w *watchdog.Watchdog) {
	i := 0
	for {
		select {
		case <-w.GetKickChannel():
			fmt.Println("kick!")
			i++
			if i == 3 {
				// stop the watchdog permanently
				w.Stop()
				return
			}
		}
	}
}(wd)
```
Reset the watchdog timer:
```go
w.Reset()
```

## Online example
[Go Playground](https://play.golang.org/p/wLtW6V4k8k7)
