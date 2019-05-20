package main

import (
	"context"
	"fmt"
	"time"

	"github.com/michaeldistler/gomemcache/memcache"
	"golang.org/x/sync/semaphore"
)

//Test Empties the caches and then checks multiple times if it is empty
func main() {
	ctx := context.TODO()
	conc := 250
	sem := semaphore.NewWeighted(int64(conc))
	responseChecker := []bool{}

	client := memcache.New("def", "10.0.6.130:11211", "10.0.7.165:11211", "10.0.7.77:11211")
	client.Timeout = 60 * time.Second
	client.FlushAll()

	i := 0
	for i < 500 {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		go func(i int) {
			defer sem.Release(1)
			response, err := client.IsACacheEmpty()
			if err != nil {
				fmt.Println(err)
			}
			responseChecker = append(responseChecker, response)

		}(i)
		i++
	}
	if err := sem.Acquire(ctx, int64(conc)); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v", err)
	}

	for _, response := range responseChecker {
		if !response {
			fmt.Println("Test Failed, a response came back FALSE")
			return
		}
	}
	fmt.Println("Test Passed all responses came back TRUE")
	return

}
