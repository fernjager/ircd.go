package main

import (
	"bufio"
	"net"
	"sync/atomic"
    "math/rand"
	"time"
    "strconv"
)

const (
	MSG_RAW     = 0
	MSG_COMMAND = 1
    MSG_PING = 2
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
    connected int32
	identified int32
    lagtime int32
    ping_num int32

}

func InitUser(c net.Conn) *User {
	readbuf := bufio.NewReader(c)
	newUser := &User{c, "", "", "", "", 0, make(map[string]*Channel), "", make([]byte, 9), make(chan *Message), readbuf, 1, 0, 0, 0}

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
            print("Received :" + data)
		}
		//if err != nil {
		ParseRawMessage(u, data)

        // reset lag timer so that they don't timeout
        atomic.StoreInt32( &u.lagtime , 0)
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
                print("Error" + err.Error())
				//u.Disconnect()
				//return
			}
        case MSG_PING:
            ping_num := rand.Int31()
            atomic.StoreInt32(&u.ping_num, ping_num)
            _, err := u.conn.Write([]byte( ":" + SERVER_NAME+" PING :" + strconv.FormatInt(int64(ping_num),10) + "\n" ))
            if err != nil {
                print("Error" + err.Error())
            }
			//case MSG_NOTIFY:
			//    _, err := u.conn.Write([]byte(MessageToRawString(u, msg))) // don't care about return value
		}

	}
}

func (u *User) Ping() {
    u.Send( &Message{MSG_PING, ""} )
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
    print("Login");
	/* If they are not logged in, then log them in */
	if atomic.LoadInt32( &u.identified ) == 0 {
		u.username = userName
		u.realname = realName

		/* Add to global datastore */
		Data.putUser(u)

		/* Let user know he/she is added */
		u.Send(&Message{MSG_RAW, (":" + SERVER_NAME + " NOTICE * :*** LOGGED IN\n")})
		atomic.StoreInt32(&u.identified, 1)
	} else {
		// **TODO** user already exists 
	}
}
func (u *User) pingTimer() {
	time.Sleep(INIT_TIMEOUT * time.Second)
	/* If they are not identified by now, then kill connection */
	if atomic.LoadInt32( &u.identified ) == 0 {
		u.Disconnect()
        return
	}
    /* Now, for the rest of the session, ping them and make sure ping timeout is not exceeded */
    for u.connected == 1{
        time.Sleep(1 * time.Second)

        currTime := atomic.LoadInt32( &u.lagtime)
        // check to see if time is over timeout, then kill
        if currTime > PING_TIMEOUT{
            u.Disconnect()
        } else if currTime > 0 && currTime % PING_INTERVAL == 0{
            u.Ping()
        }
        // Increment
        atomic.AddInt32(&u.lagtime, 1)
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
    atomic.StoreInt32(&u.connected, 0)
	// Manually call userWriteThread to ensure final messages are written out
	u.conn.Close()
}
