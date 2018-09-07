package main

import (
	"log"
	"book/queue"
)

func main()  {
	forever := make(chan bool)

	msgs := queue.PullToQueue("book_test")

	go func (){
		for d := range msgs{
			log.Printf("Received a message : %s",d.Body)
		}
	}()

	log.Printf("[*] waiting for message . To exit press CTRL+C")
	<-forever
}
