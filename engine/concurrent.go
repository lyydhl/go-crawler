package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	ConfigureMasterWorkerChan(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

/* 第二版 */
// func (e *ConcurrentEngine) Run(seed ...Request) {

// 	in := make(chan Request)
// 	out := make(chan ParserResult)
// 	e.Scheduler.ConfigureMasterWorkerChan(in)

// 	for i := 0; i < e.WorkerCount; i++ {
// 		createWorker(in, out)
// 	}

// 	for _, r := range seed {
// 		e.Scheduler.Submit(r)
// 	}

// 	itemCount := 0
// 	for {
// 		result := <-out
// 		for _, item := range result.Items {
// 			fmt.Printf("Got item #%d %v\n", itemCount, item)
// 			itemCount++
// 		}

// 		for _, request := range result.Requests {
// 			e.Scheduler.Submit(request)
// 		}
// 	}

// }

// func createWorker(in chan Request, out chan ParserResult) {
// 	go func() {
// 		for {
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }

/* 第二版 */

/* 第三版 */
func (e *ConcurrentEngine) Run(seed ...Request) {

	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seed {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item #%d %v\n", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/* 第三版 */
