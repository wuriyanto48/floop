package floop

import (
	"time"
)

//fTask type function
type fTask func()

//floop
type floop struct {
	every time.Duration
}

//New function, the floop constructor
func New(every time.Duration) *floop {
	return &floop{every: every}
}

//Start function, execute function given to parameter and stop after duration > current time
func (f *floop) Start(task fTask, duration time.Duration) {
	stop := f.exec(task)
	<-time.After(duration)
	close(stop)
}

//exec function, execute function concurrently
func (f *floop) exec(task fTask) chan bool {
	stop := make(chan bool, 1)
	tick := time.NewTicker(time.Second * f.every)
	go func() {
		for {
			select {
			case <-tick.C:
				task()
			case <-stop:
				return
			}
		}
	}()
	return stop
}
