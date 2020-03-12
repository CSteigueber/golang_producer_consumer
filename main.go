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

// consumer is returning a boolean which isn't used in program execution but eases testing
func consumer(c chan *Tweet, consumerHasFinished chan bool) (result bool) {
	tweet := <-c
	if tweet.IsTalkingAboutGo() {
		fmt.Printf("%s\ttweets about golang\n", tweet.Username)
		result = true
	} else {
		fmt.Printf("%s\tdoes not tweet about golang\n", tweet.Username)
		result = false
	}
	consumerHasFinished <- true
	return result
}

func main() {
	var start = time.Now()
	var stream = GetMockStream()
	var length = len(stream.tweets)
	c := make(chan *Tweet)
	defer close(c)
	consumerHasFinished := make(chan bool)
	// Build Producer for each tweet
	for i := 0; i < length; i++ {
		go producer(stream, c, i)
	}
	// Build Consumer for each tweet
	for i := 0; i < length; i++ {
		go consumer(c, consumerHasFinished)
	}

	// Waiting for all consumers to finish
	for i := 0; i < length; i++ {
		<-consumerHasFinished
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}
