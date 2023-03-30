package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying", err)
		time.Sleep(time.Second << uint(tries)) //指数退避策略
	}
	return fmt.Errorf("server %s failed respond after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is done: %v\n", err)
		os.Exit(1)
	}
}
