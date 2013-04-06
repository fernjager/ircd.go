package main

const (
	READ_USER  = 0
	WRITE_USER = 1

	READ_CHANNEL  = 2
	WRITE_CHANNEL = 3
)

//users := make( map[string](*Users) )
type DataStore struct {
	users    map[string](*User)
	channels map[string](*Channel)

	requests chan *DataRequest
}

type DataRequest struct {
	requestType   int8
	user          *User
	userWriteBack chan *User

	channel       *Channel
	chanWriteBack chan *Channel
}

func DataStoreInit() *DataStore {

	dataStore := &DataStore{make(map[string](*User)), make(map[string](*Channel)), make(chan *DataRequest)}
	go dataStore.dataThread()
	return dataStore
}

func (data *DataStore) dataThread() {
	for request := range data.requests {
		switch request.requestType {
		/* Read user */
		case READ_USER:
			request.userWriteBack <- data.users[request.user.nick]
		/* Write */
		case WRITE_USER:
			data.users[request.user.nick] = request.user
		}
	}
}

func (data *DataStore) getUser(user *User, writeBack (chan *User)) {
	data.requests <- &DataRequest{READ_USER, user, writeBack, nil, nil}
}

func (data *DataStore) putUser(user *User) {
	// use channels
	data.requests <- &DataRequest{WRITE_USER, user, nil, nil, nil}
	if DEBUG {
		print("Added user " + user.nick + " to the store")
	}
}

func (data *DataStore) getChannel(channel *Channel, writeBack (chan *Channel)) {
	data.requests <- &DataRequest{READ_CHANNEL, nil, nil, channel, writeBack}
}

func (data *DataStore) putChannel(channel *Channel) {
	// use channels
	data.requests <- &DataRequest{WRITE_CHANNEL, nil, nil, channel, nil}
}

func (data *DataStore) RemoveUser(nickname string) {
	if data.users[nickname] != nil {
		delete(data.users, nickname)
	}

}

// remove user, use channel <-
//
