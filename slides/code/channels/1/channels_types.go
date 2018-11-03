package main

type result struct {
	value string
	err   error
}

// Only a sender can close a channel...

// Send, receive, and close
type sendReceiver chan result

// Only send and close
type sender chan<- result

// Only receive, but not close
type receiver <-chan result
