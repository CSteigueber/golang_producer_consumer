# golang_producer_consumer
My solution to the golang producer / consumer challenge

To achieve a faster performance I opened a go routine for each tweet. To make sure all go routines fully execute before program termination each consumer is sending on a boolean channel 'consumerHasFnished' after printing its output . At the end of the main program, a loop is iterating over a statement receiving from consumerHasFinished for each consumer that was created.

# Performance
I tested performance with an increased sample size. The program can analyse about 3000 tweets in ca 680ms.


# Testing
Unit testing covers ca 63% of the whole package and 100% of the file mockstream.go.
