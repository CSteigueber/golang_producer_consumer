# golang_producer_consumer
My solution to the golang producer / consumer challenge

To achieve a faster performance I opened a channel and a go routine for each tweet. Do make sure all go routines fully execute before program termination I chose to let the receiver (consumer) close the channel, which normally should not be done. However, in this case I decided to make an exception to avoid using global variables for checking if all channels are closed.

# Bug
When many tweets are checked (>2000) some tweets are skipped. =>channels closed before producer can send imput => channel closing back to the producer, use global variable to check if all consumers are done...

Unit testing only covers 13% of the code and tests only the function "IsTalkingAboutGo" because this is the only function with a return statement.
