package retry

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoNotRetryWhenSuccessed(t *testing.T) {
	r, err := Retry(
		10,
		func(n int, r interface{}) time.Duration {
			return 0 * time.Second
		},
		func() (interface{}, error) {
			return "successed", nil
		},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, r, "successed")
	}
}

func TestRetry10Times(t *testing.T) {
	c := 0
	_, err := Retry(
		10,
		func(n int, r interface{}) time.Duration {
			return 0 * time.Second
		},
		func() (interface{}, error) {
			c++
			return nil, fmt.Errorf("dummy")
		},
	)
	assert.Error(t, err)
	assert.Equal(t, c, 10)
}
