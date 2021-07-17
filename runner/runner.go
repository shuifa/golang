package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish task within the timeout")
	ErrInterrupt = errors.New("receive interrupt from os")
)

type Runner struct {
	Interrupt chan os.Signal
	complete  chan error
	Timeout   <-chan time.Time
	Tasks     []func(int)
}

func New(t time.Duration) *Runner {
	return &Runner{
		Interrupt: make(chan os.Signal),
		complete:  make(chan error),
		Timeout:   time.After(t),
		Tasks:     make([]func(int), 0),
	}
}

func (r *Runner) AddTasks(tasks ...func(int)) {
	r.Tasks = append(r.Tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.Tasks {
		select {
		case <-r.Interrupt:
			signal.Stop(r.Interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}
func (r *Runner) Start() error {
	signal.Notify(r.Interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.Timeout:
		return ErrTimeout
	}
}
