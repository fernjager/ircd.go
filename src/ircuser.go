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
    
    /* channels */
    channels map[string] *Channel

    /* Statistics */
    lastseen string
    modes []byte

    out chan *Message
    readbuf *bufio.Reader
}

func InitUser( c net.Conn ) *User{
    readbuf := bufio.NewReader(c)
    newUser := &User{ c, "", "", "", "", 0, make(map[string] *Channel), "",  make([]byte, 9), make( chan *Message ), readbuf }

    go newUser.userReadThread();
    go newUser.userWriteThread();

    return newUser
}

func(u *User) userReadThread(){
    for {
        data, _ := u.readbuf.ReadString('\n')
        ParseRawMessage( u, data )
    }
}

func(u *User) userWriteThread(){
    for msg := range u.out {
        u.conn.Write( []byte(MessageToRawString(u,msg))) // don't care about return value
    }
}

func(u *User) Send( data *Message ) {
    u.out <- data
}

/* Handle user login routine */
func(u *User) Login( userName string, realName string ){
    /* If username and realname is not set, then log them in */
    if u.username == "" && u.realname == "" {
        u.username = userName
        u.realname = realName

        /* Add to global datastore */
        Data.putUser( u.username, u )
    }
}

func(u *User) ChangeNick( newNick string ){
    u.nick = newNick
    // make an announcement in all channels user is connected to
    if u.username=="" && u.realname==""{

    }
}

func(u *User) Disconnect(){
    // Manually call userWriteThread to ensure final messages are written out

    u.conn.Close()
}