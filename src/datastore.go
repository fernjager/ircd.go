package main


//users := make( map[string](*Users) )
type DataStore struct{
    users map[string](*User)
    channels map[string](*Channel)

}

func DataStoreInit() *DataStore{
    return &DataStore{ make(map[string](*User)), make(map[string](*Channel)) }
}

func(data *DataStore) getUser( username string ){

}

func(data *DataStore) putUser( username string, user *User ){
    data.users[username] = user
}

// remove user, use channel <-
//