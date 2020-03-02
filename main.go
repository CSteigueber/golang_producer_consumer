//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, c chan *Tweet, start int) {
	stream.pos = start
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		c <- tweet
	}
	close(c)
}

func consumer(c chan *Tweet) {
	for {
		tweet, open := <-c
		if !open {
			break
		}

		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	//length:= len(stream.tweets)
	c1 := make(chan *Tweet)
	c2 := make(chan *Tweet)
	c3 := make(chan *Tweet)
	c4 := make(chan *Tweet)
	c5 := make(chan *Tweet)

	// Producer
	go producer(stream, c1, 0)
	go producer(stream, c2, 1)
	go producer(stream, c3, 2)
	go producer(stream, c4, 3)
	go producer(stream, c5, 4)

	// Consumer
	go consumer(c1)
	go consumer(c2)
	go consumer(c3)
	go consumer(c4)
	consumer(c5)

	fmt.Printf("Process took %s\n", time.Since(start))
}
