# golang_producer_consumer
My solution to the golang producer / consumer challenge

To achieve a faster performance I opened a channel and a go routine for each tweet. To make sure all go routines fully execute before program termination each consumer is increasing the global variable 'counter' by one ('counter++') after printing its output . At the end of the program, an infite loop is checking if counter equals length (amount of tweets in input) and broken if so.

# Performance
I tested performance with an increased sample size. The program can analyse about 3000 tweets in ca 680ms.


# Testing
Unit testing covers ca 72% of the whole package and 100% of the file mockstream.go.
