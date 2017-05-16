package retry

import "time"

// ExponentialBackOff Interval
var ExponentialBackOff = func(n int, result interface{}) time.Duration {
	return time.Duration(n) * time.Duration(n) * time.Second
}

// Retry run your func "fn", if it failed, Retry will retry "n" times with "interval" algorithm
//     retry.Retry(10, retry.ExponentialBackOff, func() (interface{}, error) {...})
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
