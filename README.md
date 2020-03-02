# golang_producer_consumer
My solution to the golang producer / consumer challenge

To achieve a faster performance I opened a channel and a go routine for each tweet. To make sure all go routines fully execute before program termination each consumer is increasing the global variable 'counter' by one ('counter++') after printing its output . At the end of the program, an infite loop is checking if counter equals length (amount of tweets in input) and broken if so.

# Bug => fixed
When many tweets are checked (>2000) some tweets are skipped. =>channels closed before producer can send input => channel closing back to be done by the producer, use global variable to check if all consumers are done...


# Testing
Unit testing only covers 13% of the code and tests only the function "IsTalkingAboutGo" because this is the only function with a return statement.
