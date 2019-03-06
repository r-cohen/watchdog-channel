package whatchdog

import "time"

// Watchdog ... use the GetKickChannel() to get notified when the watchdog barks
type Watchdog struct {
	timer *time.Timer
	timeout time.Duration
	kickChan chan bool
	stopChan chan bool
}

// NewWatchdog ... ctor
func NewWatchdog(timeout time.Duration) *Watchdog {
	w := &Watchdog{
		timer: time.NewTimer(timeout),
		timeout: timeout,
		kickChan: make(chan bool, 1),
		stopChan: make(chan bool, 1),
	}
	w.timerExpireWorker()
	return w
}

func (w *Watchdog) timerExpireWorker() {
	go func(w *Watchdog) {
		for {
			select {
			case <- w.stopChan:
				return
			case <- w.timer.C:
				w.kickChan <- true
				w.Reset()
			}
		}	
	}(w)
}

// GetKickChannel ... returns the kick channel
func (w *Watchdog) GetKickChannel() chan bool {
	return w.kickChan
}

// Stop ... stops the watchdog
func (w *Watchdog) Stop() {
	w.timer.Stop()
	w.stopChan <- true
}

// Reset ... resets the timer
func (w *Watchdog) Reset() {
	w.timer.Stop()
	w.timer = time.NewTimer(w.timeout)
}
