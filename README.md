# golang_producer_consumer
My solution to the golang producer / consumer challenge

To achieve a faster performance I opened a channel and a go routine for each tweet. To make sure all go routines fully execute before program termination each consumer is increasing the global variable 'counter' by one ('counter++') after printing its output . At the end of the program, an infite loop is checking if counter equals length (amount of tweets in input) and broken if so.


# Testing
Unit testing covers ca 29% for now of the whole package and 100% of the file mockstream.go.
I changed the function func GetMockStream into func GetStream in the file mockstream.go to ease testing this function.
