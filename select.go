package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()
	/*for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}*/
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg1 := <-c2:
			fmt.Println(msg1)
		}
	}

}
