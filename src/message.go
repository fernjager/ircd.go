package main

import (
	"strings"
)

type Message struct {
	msgType int
	msgStr  string
}

func ParseRawMessage(user *User, rawString string) *Message {
	parts := strings.Split(rawString, " ")

	//if( len(parts) < 1 )

	command := strings.ToUpper(parts[0])
	switch command {
	case "NICK":
		user.ChangeNick(parts[1])
	case "USER":
		user.Login(parts[1], strings.Join(parts[4:], " ")[1:])

		//default:

	}

	return &Message{MSG_COMMAND, rawString}
}

func MessageToRawString(user *User, msg *Message) string {
	Conf.getOption("hostname")
	return msg.msgStr
}

func ServerNotice(u *User, message string) {
	//u.Send(NOTICE, totalMessage);
	// send messge to everyone if user is nil
	/* if( u == nil ){
		for nickname,userObj := range Data{

		}
	} */
}
