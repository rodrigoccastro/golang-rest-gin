package main

import "time"

type Client struct {
	id      string
	name    string
	email   string
	phone   string
	address string
	created time.Time
	updated time.Time
}