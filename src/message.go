package main

type Message struct{
    msgType string
    msgStr string
}

func ParseRawMessage( rawString string ) *Message{
    return &Message{ "PRIVMSG", rawString }
}