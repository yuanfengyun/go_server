package main

type Account struct {
	account string `json:"account"`
	passwd  string
	level   uint32
	score   uint64
}
