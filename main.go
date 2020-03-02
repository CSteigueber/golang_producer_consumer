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
	tweet, _ := stream.Next()
	c <- tweet

}

func consumer(c chan *Tweet) {
	tweet := <-c

	if tweet.IsTalkingAboutGo() {
		fmt.Println(tweet.Username, "\ttweets about golang")
	} else {
		fmt.Println(tweet.Username, "\tdoes not tweet about golang")
	}

	close(c)
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	length := len(stream.tweets)
	var c []chan *Tweet
	for i := 0; i < length; i++ {
		c = append(c, make(chan *Tweet))
	}
	// Producer

	for i := 0; i < length; i++ {
		go producer(stream, c[i], i)
	}

	// Consumer
	for i := 0; i < length-1; i++ {
		go consumer(c[i])
	}
	consumer(c[length-1])
	for {
		allClosed := true
		for i := 0; i < length; i++ {
			_, open := <-c[i]
			if open {
				allClosed = false
			}
		}
		if allClosed {
			break
		}
	}

	fmt.Printf("%d tweets analysed\n", length)
	fmt.Printf("Process took %s\n", time.Since(start))
}
