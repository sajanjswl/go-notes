package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Concurrency vs Paralleism
// * If you want to run goroutines in parallel, Go provides the ability to add more via the GOMAXPROCS environment variable or runtime function
// * Concurrency means when a processor is shared by multiple thread/goroutines
// * Parallelism is when two or more therads are executing code simultaneously against different processor
// * If you configure the runtime to use more than one logical processor, the scheduler will distribute goroutines between these logical processors which will result in goroutines running on different operating system threads.
// *  However, to have true parallelism you need to run your program on a machine with multiple physical processors
// * Channels are the way in Go we write safe and elegant concurrent programs that eliminate race conditions and make writing concurrent programs fun again.

// Paralleism example when actual dual core processor is available
func main() {
	// channelExample()
	// selectExample()
	// PrintForTenSeconds()
	PrintAtInterfal()

}

func PrintAtInterfal() {

	ticker1 := time.NewTicker(10 * time.Millisecond)
	done := make(chan bool)

	go func() {

		for {

			select {
			case t := <-ticker1.C:
				fmt.Println("ticker ", t)
			case <-done:
				return

			}
		}

	}()

	time.Sleep(50 * time.Millisecond)
	ticker1.Stop()
	done <- true

}

func PrintForTenSeconds() {
	timer1 := time.NewTimer(1 * time.Millisecond)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {

			select {

			case <-timer1.C:
				fmt.Println("Time out")
				return
			default:
				fmt.Println("default", i)
			}

		}
	}()

	wg.Wait()

}

func channelExample() {

	ch := make(chan string)

	go func() {
		ch <- "Heloo"
	}()

	fmt.Println(<-ch)

}

func selectExample() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Nanosecond)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(3 * time.Nanosecond)
		c2 <- "two"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(2 * time.Nanosecond):
		fmt.Println("time out")
	}

	//  j, ok  := <-jobs   but for map we have ok,v:=map[key]

}

func goroutineExample() {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
