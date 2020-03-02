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
		}
	}
}
