package retry

import "time"

var exponentialBackOff = func(n int) time.Duration {
	return time.Duration(n) * time.Duration(n) * time.Second
}

func Retry(
	n int,
	interval func(n int, result interface{}) time.Duration,
	fn func() (interface{}, error),
) (interface{}, error) {
	return loop(1, n, interval, fn)
}

func loop(
	i, n int,
	interval func(n int, result interface{}) time.Duration,
	fn func() (interface{}, error),
) (interface{}, error) {
	res, err := fn()
	if i >= n {
		return res, err
	}

	if err != nil {
		time.Sleep(interval(i, res))
		return loop(i+1, n, interval, fn)
	}

	return res, nil
}
