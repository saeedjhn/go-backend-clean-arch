package main

type SMSSender interface {
	Send(destination, message string) error
}
