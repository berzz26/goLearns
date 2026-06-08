package main
//  imagine channels as a shared pipeline between go routines. go routines send/recieve data from and to channels
import (
	"fmt"
	// "sync"
)
func worker(ch chan int) {
	ch <- 2

}
func main() {
	//intialize a channel ch that transports integers
	ch := make(chan int)
	//initialize a worker that takes a channel and operates on it
	//realise that there is no wg used here. why because channels are syncronized

	go worker(ch)
	//worker sends int 2 to the channel
	//mamin recieved the value here in msg
	//this thing, msg := <-ch waits automatically and hatls the go routine until msg is recieved thats why no wgs
	msg := <-ch
	fmt.Print(msg)

	//deadlock eg
	// ch := make(chan int)

	// ch <- 42

	// what happened here? no body recieved so sender waiting but reciever doesnt exists

}
