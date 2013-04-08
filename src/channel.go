package main

type Channel struct {
	name           string
	topic          string
	modes          *Mode
	channelMembers map[string]*User
}
