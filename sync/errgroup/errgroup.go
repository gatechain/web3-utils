package errgroup

import (
	"strings"
	"sync"
)

type Errors []error

func (errors Errors) Error() string {
	if len(errors) == 0 {
		return ""
	}
	msg := make([]string, 0, len(errors))
	for _, err := range errors {
		msg = append(msg, err.Error())
	}
	return strings.Join(msg, "\n")
}

type Group struct {
	wg     sync.WaitGroup
	mu     sync.Mutex
	errors Errors
}

func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.mu.Lock()
			defer g.mu.Unlock()
			g.errors = append(g.errors, err)
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if len(g.errors) == 0 {
		return nil
	}
	errors := g.errors
	g.errors = nil
	return errors
}
