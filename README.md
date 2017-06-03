# retry

[![Build Status](https://travis-ci.org/ara-ta3/retry.svg?branch=master)](https://travis-ci.org/ara-ta3/retry)

# Example

main.go

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ara-ta3/retry"
)

func main() {
	c := someClient{
		n: 0,
	}
	// retry.Retry returns the return values of "func" at third args
	r, err := retry.Retry(
		10,
		func(n int, r interface{}) time.Duration {
			// waiting time
			return 0 * time.Second
		},
		func() (interface{}, error) {
			// your codes
			// if this returns error, this will retry this "func"
			return c.DoSomething()
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	res, ok := r.(*result)

	if !ok {
		log.Fatalf("result cannot be cast to Result struct. r: %+v", r)
	}

	fmt.Printf("%+v", res)
}

type someClient struct {
	n int
}

type result struct {
	Message string
}

func (c *someClient) DoSomething() (*result, error) {
	if c.n >= 5 {
		return &result{
			Message: fmt.Sprintf("some message. n: %d", c.n),
		}, nil
	}
	c.n++
	fmt.Printf("count: %d\n", c.n)
	return nil, fmt.Errorf("return error for retry")
}
```

```zsh
$go run main.go
count: 1
count: 2
count: 3
count: 4
count: 5
&{Message:some message. n: 5}
```
