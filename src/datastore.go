package main

//users := make( map[string](*Users) )
type DataStore struct {
	users    map[string](*User)
	channels map[string](*Channel)
}

func DataStoreInit() *DataStore {
	return &DataStore{make(map[string](*User)), make(map[string](*Channel))}
}

func (data *DataStore) getUser(username string) *User {
	return data.users[username]
}

func (data *DataStore) putUser(username string, user *User) {
	data.users[username] = user
}

func (data *DataStore) RemoveUser(username string) {
	if data.users[username] != nil {
		delete(data.users, username)
	}

}

// remove user, use channel <-
//
