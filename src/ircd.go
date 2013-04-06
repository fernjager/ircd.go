package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	//"time"
)

const homePage = `<!DOCTYPE html>
<html>
<head>
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script>
</head>
<body>
    <form action="/target" id="postToGoHandler">
    <input type="submit" value="Post" />
    </form>
    <div id="result"></div>
<script>
$("#postToGoHandler").submit(function(event) {
    event.preventDefault();
    $.post("/target", JSON.stringify({"Param1": "Value1"}),
        function(data) {
            $("#result").empty().append(data);
        }
    );
});
</script>
</body>
</html>`

var Conf *Config
var Data *DataStore

var DEBUG = true

const initTimeout = 5    // disconnect connection if we don't receive anything in 5 seconds
const pingTimeout = 5    // if they don't respond after 5 ping messages, they are disconnected
const PING_INTERVAL = 30 // seconds between pings
const SERVER_NAME = "irc.test.net"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"leads\":\"1\",\"success\":true}")
}

func target(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		fmt.Fprintf(w, "Couldn't read request body: %s", err)
	} else {
		dec := json.NewDecoder(strings.NewReader(string(body)))
		var m Message
		if err := dec.Decode(&m); err != nil {
			fmt.Fprintf(w, "Couldn't decode JSON: %s", err)
		} else {

		}
	}
}

func startWeb() {
	http.HandleFunc("/", home)
	http.HandleFunc("/target", target)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func startLinks() {

}

func handleConnection(c net.Conn) {
	InitUser(c)

	// Now get NICK
	//daytime := time.Now().String()

	//user.Disconnect();
}

func startIRCd() {
	ln, err := net.Listen("tcp", ":6667")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		print("connection from " + conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func main() {
	Conf = ConfigInit()
	Data = DataStoreInit()
	go startWeb()
	go startLinks()
	startIRCd()
}
