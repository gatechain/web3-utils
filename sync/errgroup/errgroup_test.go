package errgroup

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	var g Group
	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		return errors.New("hi")
	})
	g.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return errors.New("hello")
	})
	err := g.Wait()
	assert.Error(t, err)
	assert.Equal(t, "hello\nhi", err.Error())

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	g.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	})
	err = g.Wait()
	assert.NoError(t, err)
}
