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

// global variables
var counter = 0

func producer(stream Stream, c chan *Tweet, start int) {
	stream.pos = start
	tweet, _ := stream.Next()
	c <- tweet
	close(c)
}

// consumer is returning a bolean which isn't used in program execution but eases testing
func consumer(c chan *Tweet) (result bool) {
	tweet := <-c
	if tweet.IsTalkingAboutGo() {
		fmt.Printf("%s\ttweets about golang\n", tweet.Username)
		result = true
	} else {
		fmt.Printf("%s\tdoes not tweet about golang\n", tweet.Username)
		result = false
	}
	counter++
	return result
}
func makeChannels(length int) (c []chan *Tweet) {
	for i := 0; i < length; i++ {
		c = append(c, make(chan *Tweet))
	}
	return c
}
func main() {
	var start = time.Now()
	var stream = GetMockStream()
	var length = len(stream.tweets)
	var c = makeChannels(length)
	// Producer
	for i := 0; i < length; i++ {
		go producer(stream, c[i], i)
	}
	// Consumer
	for i := 0; i < length; i++ {
		go consumer(c[i])
	}

	// Waiting for the consumers to finish
	for counter != length {
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}
