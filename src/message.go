package main

import (
    "strings"
)
type Message struct{
    msgType string
    msgStr string
}

func ParseRawMessage( user *User, rawString string ) *Message{
    parts := strings.Split( rawString, " " )
    
    //if( len(parts) < 1 )

    command := strings.ToUpper( parts[0] )
    switch command {
        case "NICK":
            user.ChangeNick( parts[1] )
        case "USER":
            user.Login( parts[1], strings.Join( parts[4:], " " )[1:])

        //default:


    }

    return &Message{ command, rawString }
}


func MessageToRawString( user *User, msg *Message ) string{
    Conf.getOption( "hostname" )
    return msg.msgStr
}