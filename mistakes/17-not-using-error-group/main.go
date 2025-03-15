package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
	}

	res, err := handler(ctx, urls)
	if err != nil {
		fmt.Printf("errr => %#v \n", err.Error())
		cancel()
		return
	}
	fmt.Println(res)
}

func handler(ctx context.Context, urls []string) ([]string, error) {
	result := make([]string, len(urls))

	g, ctx := errgroup.WithContext(ctx)

	for i, url := range urls {
		g.Go(func() error {
			timeOut := i * 3
			if timeOut == 0 {
				timeOut = 1
			}

			fmt.Printf("I: %d URL:%s \n", i, url)

			r, err := queryUrl(ctx, url, timeOut)
			if err != nil {
				return err
			}
			result[i] = r
			return nil
		})
	}

	return result, g.Wait()
}

func queryUrl(ctx context.Context, url string, timeOut int) (string, error) {
	fmt.Printf("Waiting %d seconds for: %s \n", timeOut, url)
	time.Sleep(time.Second * time.Duration(timeOut))
	ctx.Done()
	return "", fmt.Errorf("timeout after %d seconds for: %s", timeOut, url)
}
