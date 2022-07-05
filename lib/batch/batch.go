package batch

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, n)

	eg, _ := errgroup.WithContext(context.Background())
	eg.SetLimit(int(pool))

	var mu sync.Mutex
	var i int64

	for i = 0; i < n; i++ {
		id := i

		eg.Go(func() error {
			user := getOne(id)

			mu.Lock()
			defer mu.Unlock()

			res[id] = user

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil
	}

	return res
}
