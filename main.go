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
	close(c)

}

// consumer is returning a bolean which isn't used in program execution but eases testing
func consumer(c chan *Tweet) bool {
	tweet := <-c
	var result bool
	if tweet.IsTalkingAboutGo() {
		fmt.Println(tweet.Username + "\ttweets about golang")
		result = true
	} else {
		fmt.Println(tweet.Username + "\tdoes not tweet about golang")
		result = false
	}
	counter++
	return result

}
func makeChannels(length int) []chan *Tweet {
	var c []chan *Tweet
	for i := 0; i < length; i++ {
		c = append(c, make(chan *Tweet))
	}
	return c
}

// global variable
var counter = 0

func main() {

	counter = 0
	start := time.Now()
	stream := GetMockStream()
	length := len(stream.tweets)
	c := makeChannels(length)
	// Producer

	for i := 0; i < length; i++ {
		go producer(stream, c[i], i)
	}
	// Consumer
	for i := 0; i < length-1; i++ {
		go consumer(c[i])
	}
	consumer(c[length-1])

	// Waiting for the consumers to finish
	for {
		if counter == length {
			break
		}
	}

	fmt.Printf("Process took %s\n", time.Since(start))
}
