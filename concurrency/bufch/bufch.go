package main

import "fmt"

// this demonstrates both buffered channel,
// as well as select behavior: it doesn't block if default branch allow it to proceed.

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	// now we read it back
	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to to done")
			//break  ... this is interesting, the break doesn't break out of loop here.
			// is break in select only break out select {}?
		}
	}
}
