package main

import "testing"

func TestGetMockStream(t *testing.T) {
	stream := GetMockStream()
	for i := 0; i < len(mockdata); i++ {

		if stream.tweets[i] != mockdata[i] {
			t.Errorf("Error! Tweet was not loaded correctly! Expected '%s' but received '%s'", mockdata[i].Text, stream.tweets[i].Text)
		}

	}

}

func TestNext(t *testing.T) {
	tables := []Stream{
		{1, []Tweet{Tweet{"user", "tweet"}, {"user2", "tweet2"}}},
		{0, []Tweet{}},
		{0, []Tweet{Tweet{"Han Solo", "Am I alone in this Stream? Wher is Chewbaka?"}}},
		{1, []Tweet{Tweet{"Trollwurf", "Husaah it works!"}, {"Claas", "Who is that Trollwurf-guy?"}, {"Tester69", "Cheap online test? Click here!"}}},
	}
	for _, table := range tables {
		start := table.pos
		if table.pos >= len(table.tweets) {
			_, err := table.Next()
			if err != ErrEOF {
				t.Errorf("Error! Missed end of file!")
			}
		} else {
			tweet, _ := table.Next()
			if *tweet != table.tweets[start] {
				t.Errorf("Error! wrong tweet loaded!")
			}
			if start == table.pos {
				t.Errorf("Error! Position in stream was not increased!")
			}
		}
	}
}
func TestMakeChannels(t *testing.T) {
	tables := []int{
		1, 2, 5,
	}
	for _, table := range tables {
		c := makeChannels(table)
		if len(c) != table {
			t.Errorf("Error! Wrong amount of chanels produced! Expected %v and got %v", table, len(c))
		}
	}
}

func TestProducer(t *testing.T) {
	tables := []struct {
		stream Stream
		c      chan *Tweet
		start  int
		result Tweet
	}{
		{Stream{0, []Tweet{Tweet{"user", "tweet"}, {"user", "tweet2"}}}, make(chan *Tweet, 1), 0, Tweet{"user", "tweet"}},
		{Stream{0, []Tweet{Tweet{"user", "tweet"}, {"user", "tweet2"}, {"user2", "tweet3"}}}, make(chan *Tweet, 1), 1, Tweet{"user", "tweet2"}},
	}
	for _, table := range tables {
		producer(table.stream, table.c, table.start)

		tweet := <-table.c
		if tweet.Username != table.result.Username {
			t.Errorf("Error! Producer produced wrong username! Expected %s but got %s", table.result.Username, tweet.Username)
		}
		if tweet.Text != table.result.Text {
			t.Errorf("Error! Producer produced wrong tweet! Expected %s but got %s", table.result.Text, tweet.Text)
		}
	}
}
func TestConsumer(t *testing.T) {
	tables := []struct {
		tweet  Tweet
		result bool
	}{
		{Tweet{"Trollwurf", "golang is so awesome"}, true},
		{Tweet{"SomeRandomDude", "qddfgolangdcenjen"}, true},
		{Tweet{"Pumuckel", " dont forget about the gopher"}, true},
		{Tweet{"GuyFawkes", "Go get some coffee"}, false},
		{Tweet{"Tester69", "GOLANG FOR PRESIDENT!!"}, true},
		{Tweet{"Oedipus", "GOPHERS GONE WILD!!"}, true},
		{Tweet{"Neo", "12*%c!@"}, false},
	}
	c := make(chan *Tweet, 1)
	for _, table := range tables {
		c <- &table.tweet
		output := consumer(c)
		if output != table.result {
			t.Errorf("Error! Expected'%t'\tbut received\t'%t'\tfrom function consumer() with tweet %s", table.result, output, table.tweet.Text)
		}

	}
}
