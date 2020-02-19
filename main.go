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

func producer(stream Stream, c chan []*Tweet) {
	var tweets []*Tweet
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(c)
			break
		}
		//fmt.Println(tweet.Username, tweet.Text)
		tweets = append(tweets, tweet)
		c <- tweets
	}
}

func consumer(c chan []*Tweet) {
	i := 0
	for {
		tweets, open := <-c
		if !open {
			break
		}
		t := tweets[i]
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
		i++
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	c := make(chan []*Tweet, 5)
	//var tweets []*Tweet

	// Producer
	go func() {
		producer(stream, c)

	}()

	// Consumer
	consumer(c)

	fmt.Printf("Process took %s\n", time.Since(start))
}
