package main

import (
    "net"
    "bufio"
    );

// User objects
type User struct{
    conn net.Conn

    address string
    nick string
    username string
    realname string
    usermode int8
    
    /* Statistics */
    lastseen string
    modes []byte

    in chan *Message
    out chan *Message
    readbuf *bufio.Reader
}

func InitUser( c net.Conn ) *User{
    readbuf := bufio.NewReader(c)
    newUser := &User{ c, "", "", "", "", 0, "",  make([]byte, 9), make( chan *Message ), make( chan *Message ), readbuf }

    go newUser.userReadThread();
    go newUser.userWriteThread();
    return newUser
}

func(u *User) userReadThread(){
    for {
        data, _ := u.readbuf.ReadString('\n')
        u.in <- ParseRawMessage( data )
    }
}
func(u *User) userWriteThread(){
    for msg := range u.out {
        u.conn.Write( []byte(msg.msgStr)) // don't care about return value
    }
}

func(u *User) Send( data *Message ) {
    u.out <- data
}
func(u *User) Receive() *Message {
    return <- u.in
}

func(u *User) Disconnect(){
    // Manually call userWriteThread to ensure final messages are written out

    u.conn.Close()
}