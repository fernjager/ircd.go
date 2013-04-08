package main

import (
	"strings"
	"sync/atomic"
	"strconv"
)

type Message struct {
	msgType int
	msgStr  string
}

func ParseRawMessage(user *User, rawString string) *Message {
	parts := strings.Split(strings.Trim(rawString, "\n\r "), " ")

	//if( len(parts) < 1 )

	command := strings.ToUpper(parts[0])
	switch command {
	case "NICK":
		user.ChangeNick(parts[1])
	case "USER":
		user.Login(parts[1], strings.Join(parts[4:], " ")[1:])
	case "PONG":
		pong_num,err := strconv.Atoi(parts[1][1:])
		if err == nil && int32(pong_num) == atomic.LoadInt32(&user.ping_num) {
			// reset user lagtime
			atomic.StoreInt32(&user.lagtime, 0)
		}
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
