package main

import (
	"fmt"
	"my_book/workpools"
	"time"
)

var i = func() {
	fmt.Println(1)
}

func main() {
	jobs := &workpools.Jobs{}

	jobs.Job = make(chan func())

	jobs.StartPools(1)

	go func() {
		for j := 0; j < 10; j ++ {
			jobs.Job <- i
		}
		close(jobs.Job)
	}()

	time.Sleep(time.Second * 1)

	fmt.Println(jobs)

	//for v := range jobs.Result {
	//	fmt.Println(v)
	//}

}
