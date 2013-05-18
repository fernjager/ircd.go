package main

import (
    "strconv"
    "strings"
    "sync/atomic"
)

type Message struct {
    msgType int
    msgStr  string
}

func ParseRawMessage(user *User, rawString string) *Message {
    parts := strings.Split(strings.Trim(rawString, "\n\r "), " ")

    if len(parts) > 0 {
        // Determine command format i.e. :command or command
        var command string
        commandString := strings.ToUpper(parts[0])

        if commandString[0] != ':' {
            command = commandString
        } else {
            command = commandString[1:]
        }

        switch command {
        //case "PASS":

        case "NICK":
            user.ChangeNick(parts[1])
        case "USER":
            user.Login(parts[1], strings.Join(parts[4:], " ")[1:])
        case "PONG":
            pong_num, err := strconv.Atoi(parts[1][1:])
            if err == nil && int32(pong_num) == atomic.LoadInt32(&user.ping_num) {
                // Client must respond to proper pings in order to reset user lagtime
                atomic.StoreInt32(&user.lagtime, 0)
            }

            //default:

        }
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
