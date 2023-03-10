package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	//为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//最后，我们收集所有这些任务的返回值。 这也确保了所有的 worker 协程都已完成。 另一个等待多个协程的方法是使用WaitGroup。
	for a := 1; a < numJobs; a++ {
		<-results
	}
}
