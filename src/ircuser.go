package main

import (
	"bufio"
	"net"
	"sync/atomic"
	"time"
)

const (
	MSG_RAW     = 0
	MSG_COMMAND = 1
)

// User objects
type User struct {
	conn net.Conn

	address  string
	nick     string
	username string
	realname string
	usermode int8

	/* channels */
	channels map[string]*Channel

	/* Statistics */
	lastseen string
	modes    []byte

	out     chan *Message
	readbuf *bufio.Reader
	/* state */
	identified int32
}

func InitUser(c net.Conn) *User {
	readbuf := bufio.NewReader(c)
	newUser := &User{c, "", "", "", "", 0, make(map[string]*Channel), "", make([]byte, 9), make(chan *Message), readbuf, 0}

	go newUser.userReadThread()
	go newUser.userWriteThread()

	go newUser.pingTimer()

	return newUser
}

func (u *User) userReadThread() {
	for {
		data, err := u.readbuf.ReadString('\n')
		if err != nil {
			u.Disconnect()
			return
		}
		if DEBUG {

		}
		//if err != nil {
		ParseRawMessage(u, data)
		//} else {
		//Data.RemoveUser(u.nick)
		//u.Disconnect()
		//}
	}
}

func (u *User) userWriteThread() {
	for msg := range u.out {
		switch msg.msgType {
		case MSG_RAW:
			_, err := u.conn.Write([]byte(msg.msgStr)) // don't care about return value
			if err != nil {
				u.Disconnect()
				return
			}
			//case MSG_NOTIFY:
			//    _, err := u.conn.Write([]byte(MessageToRawString(u, msg))) // don't care about return value
		}

	}
}

func (u *User) Send(data *Message) {
	u.out <- data
}

/* Handle user login routine */
func (u *User) Login(userName string, realName string) {
	/* var located *User
	foundUser := make(chan *User)

	Data.getUser(u, foundUser)
	located = <-foundUser */
	/* If they are not logged in, then log them in */
	if u.identified == 0 {
		u.username = userName
		u.realname = realName

		/* Add to global datastore */
		Data.putUser(u)

		if DEBUG {
			print("Notify user")
		}
		/* Let user know he/she is added */
		u.Send(&Message{MSG_RAW, (":" + SERVER_NAME + " NOTICE * :*** LOGGED IN")})

		atomic.CompareAndSwapInt32(&u.identified, 0, 1)
	} else {
		// **TODO** user already exists 
	}
}
func (u *User) pingTimer() {
	time.Sleep(INIT_TIMEOUT * time.Second)
	/* If they are not identified by now, then kill connection */
	if u.identified == 0 {
		u.Disconnect()
	}
}
func (u *User) ChangeNick(newNick string) {
	u.nick = newNick
	// make an announcement in all channels user is connected to
	// if they are in the datastore
	//if Data.getUser(u, ) != nil {

	//}
}

func (u *User) Disconnect() {
	// Manually call userWriteThread to ensure final messages are written out
	u.conn.Close()
}
