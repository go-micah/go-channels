package main

import (
	"fmt"
	"time"
)

func longBlockingProcess(s string) {
	fmt.Printf("Starting blocking process for %v\n", s)
	time.Sleep(10 * time.Second)
	fmt.Printf("Blocking process for %v finished\n", s)
}

func longRunningProcess(s string, c chan string) {
	fmt.Printf("Starting running process for %v\n", s)
	time.Sleep(10 * time.Second)
	c <- s + ": done."
}

func main() {

	longBlockingProcess("Process A")

	go longBlockingProcess("Process B")

	c := make(chan string)

	go longRunningProcess("Process C", c)
	go longRunningProcess("Process D", c)

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("Exiting")
}
