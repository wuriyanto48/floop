package floop

import (
	"time"
)

type Floop struct {
	tick *time.Ticker
}

func New(sch time.Duration) *Floop {
	tick := time.NewTicker(time.Second * sch)
	return &Floop{tick}
}

type fTask func()

func (f *Floop) Start(task fTask) {
	for {
		<-f.exec(task)
	}
}

func (f *Floop) Stop() {
	f.tick.Stop()
}

func (f *Floop) exec(task fTask) <-chan time.Time {
	tr := make(chan time.Time)
	go func() {
		for t := range f.tick.C {
			task()
			tr <- t
		}
	}()
	return tr
}
