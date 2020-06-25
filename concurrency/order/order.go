// this program demo how to control of order of execution of go routines
// the order is defined by the channels
// in particular, closing a channel will actually "unblock" it for other goroutine who is blocking on read
// BUT: once a channel is closed; you can read from it; but you CAN't write to it

// in the following example:
// execution order: A, B, C
// where A, B can exactly be called on once, C can be called on multiple times

package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a // will block until the channel is closed
	fmt.Println("A() in the work")
	time.Sleep(time.Second)
	close(b) // will unlock whoever waiting on b channel
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B() in the work")
	time.Sleep(time.Second)
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C() in the work")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)    // C will wait on channel z
	go A(x, y) // A will wait on x, and close x at the end
	go C(z)    // call C again
	go B(y, z) // B will wait on y, and close z at the end (to unlock C)
	go C(z)    // call C again

	close(x)                    // trigger A who is waiting on x
	time.Sleep(3 * time.Second) // hang on for 3 seconds to let others do the job

}
