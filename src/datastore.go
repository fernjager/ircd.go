package main


//users := make( map[string](*Users) )
type DataStore struct{
    users map[string](*User)
    channels map[string](*Channel)

}

func DataStoreInit() *DataStore{
    return &DataStore{ make(map[string](*User)), make(map[string](*Channel)) }
}

// remove user, use channel <-
//