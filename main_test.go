package main

import "testing"

func TestIsTalkingAboutGo(t *testing.T) {
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
	for _, table := range tables {
		output := table.tweet.IsTalkingAboutGo()
		if output != table.result {
			t.Errorf("Error! Expected %t but function returned %t from user %s tweeting '%s'!!", table.result, output, table.tweet.Username, table.tweet.Text)
		} else {
			t.Logf("Success! Expected Tweet %v from %v to result in %t and got %t", table.tweet.Text, table.tweet.Username, table.result, output)
		}
	}
}
func TestGetStream(t *testing.T) {
	tables := []Stream{
		{0, []Tweet{Tweet{"user", "tweet"}, {"user2", "tweet2"}}},
	}
	for _, table := range tables {
		stream := GetStream(table.tweets)
		if stream.pos != table.pos || len(stream.tweets) != len(table.tweets) {
			t.Errorf("Error! Stream.pos was not loaded properly with data set %v", table.pos)
		}
		if len(stream.tweets) == len(table.tweets) {
			for i := 0; i < len(stream.tweets); i++ {
				if stream.tweets[i] != table.tweets[i] {
					t.Errorf("Error! Tweet %s not found in created stream", stream.tweets[i])
				}
			}
		}

	}

}
