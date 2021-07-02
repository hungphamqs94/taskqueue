package src

import (
	"math/rand"
	"time"
)

// Printer is a dummy worker that just prints received URL.
type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

// Work waits for a few seconds and print a received URL.
func (p *Printer) Work(j *Job) {
	t := time.NewTimer(time.Duration(rand.Intn(5)) * time.Second)
	defer t.Stop()
	<-t.C
}
